# Lancement et interprétation des résultats k6 🚀

## 1. Exécuter un test k6

k6 est un formidable outil qui permet d'afficher les résultats directement dans votre terminal. L'exportation en format JSON brut (`--out json=...`) est super pour des tableaux de bord (comme Grafana ou Datadog), mais elle est compliquée à lire seul.
Pour interpréter rapidement vos résultats, **inutile de regarder le JSON**, la synthèse de fin de terminal de k6 suffit amplement !

Pour lancer un de nos tests tout en affichant un résumé propre, placez-vous dans le dossier `benchlab/benchmarks/scripts` et lancez simplement le fichier :

```bash
# A. Lecture unitaire (1000 requêtes, 10 VUs)
k6 run unit-read-k6-test.js

# B. Écriture (500 requêtes, 5 VUs)
k6 run write-k6-test.js

# C. Charge progressive (10 à 100 connexions)
k6 run ramp-up-K6-test.js
```

---

## 2. Comment interpréter les résultats dans le terminal ?

À la fin de chaque test, k6 vous affichera un bloc ressemblant à ceci :

```text
     ✓ REST status is 200
     ✓ gRPC status is OK

     grpc_req_duration..............: avg=2.1ms   min=1.2ms  med=1.9ms  max=15ms   p(90)=3.4ms  p(95)=4.1ms
     http_req_duration..............: avg=4.5ms   min=3.0ms  med=4.2ms  max=22ms   p(90)=6.1ms  p(95)=8.5ms
     
     ...
     vus............................: 10
     iterations.....................: 2000
```

### Les 2 métriques les plus importantes :
Vos deux scénarios ont tourné en parallèle. k6 isole la latence REST et gRPC :
* `http_req_duration` : C'est le temps total qu'a mis votre API **REST** à répondre.
* `grpc_req_duration` : C'est le temps total qu'a mis votre API **gRPC** à répondre.

### Que signifient les chiffres ?
* **avg (moyenne)** : Le temps moyen de toutes vos requêtes combinées.
* **med (médiane)** : Si on aligne toutes les requêtes de la plus rapide à la plus lente, c'est le temps de celle du milieu. Souvent plus représentatif que la moyenne.
* **p(90) / p(95) (percentiles)** : Important ! `p(90)=3.4ms` signifie que 90% de vos requêtes ont répondu en **moins de 3.4ms**. C'est le chiffre le plus fiable pour juger de vraies performances d'un système.

### Résultat attendu 
Vous devriez clairement observer que les indicateurs `avg` et `p(95)` de la ligne **gRPC (`grpc_req_duration`)** sont inférieurs (plus rapides) à ceux de **REST (`http_req_duration`)**, surtout avec beaucoup de VUs dans le `ramp-up-K6-test.js` (car gRPC utilise le multiplexage d'HTTP/2).
