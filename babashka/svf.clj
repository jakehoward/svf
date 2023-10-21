(ns svf)

;; To Dev (emacs)
;; - `bb --nrepl-server`
;; - M-x cider-connect-clj

(require '[babashka.cli :as cli])

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
            :default "1-"}
   :help {:alias :h
          :desc "Print usage docs"
          :require false}})

(defn help []
  (println (cli/format-opts {:spec cli-opts})))

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

(defn parse-fields [fields-string]
  (cond (= "1-" fields-string)
        {:fields [0] :and-rest true}

        :else
        (let [ans
              (loop [tokens    (clojure.string/split fields-string #",")
                     z-idx-fs  []
                     and-rest  false]
                (if-let [token (first tokens)]
                  (recur (rest tokens) (into z-idx-fs (token-to-zero-idx-fields token)) (is-open? token))
                  {:fields z-idx-fs :and-rest and-rest}))]

          (if (and (seq (:fields ans))
                   (apply < (:fields ans))
                   (= (count (set (:fields ans))) (count (:fields ans)))) ;; sorted, asc
            ans

            (throw (java.lang.Exception.
                    (str "Must be strictly sorted and unique fields: " fields-string)))))))

(comment
  (let [exs ["1-" "1,2" "3,8", "9-10" "7-" "1,4-6,9-" "1,3,3" "9-8" "99,1-3,10-"]]
    (for [ex exs]
      (try
        (println "ex:" ex "parsed:" (parse-fields ex))
        (catch Exception e (println "Exception: " ex (.getMessage e))))))
  ;
  )

(defn parse-opts [{:keys [delimeter fields]}]
  {:delimeter delimeter
   :fields (parse-fields fields)})

;; bb -m svf -d ';' -f '1-'
(defn -main [& args]
  (println "args were:" args)
  (let [opts (cli/parse-opts *command-line-args* {:spec cli-opts})]

    (when (contains? opts :help)
      (help)
      (System/exit 0))

    (println "opts: " (parse-opts opts))))

