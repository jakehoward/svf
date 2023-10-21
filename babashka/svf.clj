(ns svf
  (:require [clojure.string :as str]))

;; To Dev (emacs)
;; - `bb --nrepl-server`
;; - M-x cider-connect-clj

(require '[babashka.cli :as cli])
(require '[clojure.data.csv :as csv])

(def cli-opts
  {:delimeter {:alias :d
               ;; :corece {:delimeter :string}
               :desc "The string value used to split a row into columns, defaults to comma ','"
               :require true
               :default ","}
   :fields {:alias :f
            :desc
            "The fields you want, 1 indexed split by comma, range by hyphen.
                     e.g. '1,2-5,8-' => [1,2,3,4,5,8,9...]. Defaults to '1-' (all)"
            :require true
            :coerce :string
            :default "1-"}
   :report {:alias :r
            :desc "Column count and (assumes header row) mapping of column to index"
            :require false}
   :help {:alias :h
          :desc "Print usage docs"
          :require false}})

(defn help []
  (println (cli/format-opts {:spec cli-opts})))

(defn report [head-row]
  (let [num-cols (count head-row)
        max-len  (apply max (map count head-row))
        ;; gross!
        pad-left (fn [s] (-> (java.lang.StringBuilder.)
                             (.append (str/join "" (take (- max-len (count s)) (repeat " "))))
                             (.append s)
                             .toString))]
    (println "Num columns:" num-cols)
    (println "column:idx")
    (doseq [[col idx] (map (fn [c i] [c i]) head-row (iterate inc 1))]
      (println (str (pad-left col) ": " idx)))))

(defn token-to-zero-idx-fields [token]
  (let [[lhs rhs] (clojure.string/split token #"-")]
    (if rhs
      (vec (range (dec (java.lang.Integer/parseInt lhs))
                  (java.lang.Integer/parseInt rhs)))
      [(dec (java.lang.Integer/parseInt lhs))])))

(comment
  (let [exs ["1-" "1" "8", "9-10" "7-"]]
    (for [ex exs]
      (try
        (println "ex:" ex "zero idx fields:" (token-to-zero-idx-fields ex))
        (catch Exception e (println "Exception: " ex)))))
  ;
  )

(defn is-open? [token]
  (= (last token) \-))


;; todo: currently allows 1-,2 which is a bit iffy
(defn parse-fields [fields-string]
  (cond (= "1-" fields-string)
        {:idxs [0] :and-rest true}

        :else
        (let [ans
              (loop [tokens    (clojure.string/split fields-string #",")
                     z-idx-fs  []
                     and-rest  false]
                (if-let [token (first tokens)]
                  (recur (rest tokens) (into z-idx-fs (token-to-zero-idx-fields token)) (is-open? token))
                  {:idxs z-idx-fs :and-rest and-rest}))]

          (if (and (seq (:idxs ans))
                   (apply < (:idxs ans))
                   (= (count (set (:idxs ans))) (count (:idxs ans)))) ;; sorted, asc
            ans

            (throw (java.lang.Exception.
                    (str "Must be strictly sorted and unique fields: " fields-string)))))))

(comment
  (let [exs ["1-" "1-,2" "3,8", "9-10" "7-" "1,4-6,9-" "1,3,3" "9-8" "99,1-3,10-"]]
    (for [ex exs]
      (try
        (println "ex:" ex "parsed:" (parse-fields ex))
        (catch Exception e (println "Exception: " ex (.getMessage e))))))
  ;
  )

;; slightly confusing having same name as cli/
(defn parse-opts [{:keys [delimeter fields]}]
  {:delimeter delimeter
   :fields (parse-fields fields)})

;; bb -m svf -d ';' -f '1-'
(defn -main [& args]
  (let [opts                       (cli/parse-opts *command-line-args* {:spec cli-opts})
        {:keys [delimeter fields]} (parse-opts opts)
        {:keys [idxs and-rest]}    fields
                                   ;; assumes idxs are sorted asc
        last-idx                   (last idxs)
        idxs-set                   (set idxs)
        truncate-row               (fn [row] (if-not and-rest (take (inc last-idx) row) row))
        selected-idx?              (fn [idx] (or (idxs-set idx) (and and-rest (> idx last-idx))))
                                   ;; cleaner way to filter with idx?
        filter-row                 (fn [row] (keep (fn [[idx item]]
                                                     (when (selected-idx? idx) item))
                                                   (map (fn [idx item] [idx item])
                                                        (iterate inc 0) (truncate-row row))))]

    (when (contains? opts :help)
      (help)
      (System/exit 0))

    (when (contains? opts :report)
      ;; lazy so only parses one
      (let [rows     (csv/read-csv *in* :separator (first delimeter))]
        (report (first rows))
        (System/exit 0)))

    (let [rows     (csv/read-csv *in* :separator (first delimeter))
          out-rows (->> rows
                        (map filter-row))]
      (csv/write-csv *out* out-rows :separator (first delimeter)))))

