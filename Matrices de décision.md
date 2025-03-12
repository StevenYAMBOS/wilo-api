#  Wilo

## Sommaire

- [Introduction](#introduction)
- [I - Langages Front-End](#i---langages-front-end)
- [II - Langages Back-End](#ii---langages-back-end)
- [III - Frameworks Front-End](#iii---frameworks-front-end)
- [IV - Frameworks Back-End](#iv---frameworks-back-end)
- [V - Bases de données](#v---bases-de-donnees)
- [VI - Hébergements Cloud Front-End](#vi---hebergements-cloud-front-end)
- [VII - Hébergements Cloud Back-End](#vii---hebergements-cloud-back-end)
- [VIII - Choix final](#viii---choix-final)
  - [Langage Front-End](#langage-front-end)
  - [Framework Front-End](#framework-front-end)
  - [Langage Back-End](#langage-back-end)
  - [Framework Back-End](#framework-back-end)
  - [Base de données](#base-de-donnees)
  - [Hébergement Cloud Front-End](#hebergement-cloud-front-end)
  - [Hébergement Cloud Back-End](#hebergement-cloud-back-end)

## Introduction

**Wilo - L’assistant bien-être au travail**

**Wilo** est une solution innovante conçue pour **lutter contre la sédentarité** des salariés en entreprise en intégrant **des activités physiques courtes et adaptées** directement dans leur emploi du temps. 

Grâce à une **synchronisation intelligente** avec les **agendas professionnels et personnels**, Wilo identifie les créneaux libres et propose des **mini-sessions d’activité physique** de **10 à 15 minutes** tout au long de la semaine. L’objectif est de réintroduire progressivement une routine bien-être sans perturber l’organisation quotidienne des utilisateurs.

En renseignant ton **lieu de travail et de résidence**, Wilo peut également suggérer des **activités proches** de ton environnement pour une flexibilité optimale. Selon tes préférences, tu peux également paramétrer des créneaux plus longs (**30 minutes à 1 heure**) et laisser Wilo optimiser ton emploi du temps pour te proposer **des séances adaptées**.

Wilo ne se limite pas à une expérience individuelle : tu peux choisir de pratiquer **seul ou avec tes collègues**, et même inviter d’autres personnes à se joindre à toi. La plateforme facilite aussi la synchronisation avec **l’agenda de ton/ta conjoint(e) ou de tes enfants**, permettant une organisation fluide entre vie professionnelle et personnelle.

Avec Wilo, **l’activité physique s’intègre naturellement** à ton quotidien, te permettant d’améliorer ta santé et ton bien-être sans contraintes.
## I - Langages Front-End

| Langages   | Performance | Facilité d'apprentissage | Coût de développement | Durabilité à long terme | Temps de développement | Consommation d'énergie | Popularité du langage | Maîtrise du développeur | Résultat |
| ---------- | ----------- | ------------------------ | --------------------- | ----------------------- | ---------------------- | ---------------------- | --------------------- | ----------------------- | -------- |
| JavaScript | 4           | 4                        | 4                     | 5                       | 5                      | 3                      | 5                     | 4                       | 34       |
| TypeScript | 4           | 3                        | 4                     | 5                       | 4                      | 3                      | 4                     | 3                       | 30       |
| Dart       | 3           | 3                        | 3                     | 4                       | 3                      | 4                      | 3                     | 0                       | 23       |

**Javascript**

JavaScript a initialement été conçu comme un langage côté client, il s'est ensuite étendu au développement côté serveur. Cependant, ses limites en matière de programmation orientée objet et son typage dynamique peuvent rendre son adoption complexe pour les projets de grande envergure. Pour combler ces lacunes, TypeScript a été introduit en tant que couche supplémentaire de JavaScript, en y ajoutant notamment le typage statique et d'autres fonctionnalités facilitant la gestion de projets à grande échelle.

**Typescript**

TypeScript complète Javascript, c'est un langage orienté objet qui offre des fonctionnalités absentes en JavaScript, comme le typage statique et le support des interfaces. Le typage statique permet de détecter des erreurs dès la compilation, ce qui réduit les erreurs lors des compilation. TypeScript est compatible avec JavaScript et permet d'utiliser les bibliothèques JavaScript existantes tout en intégrant du code TypeScript, donc les 2 langages se complètes.

**Source :**
- [dart-vs-javascript](https://stackshare.io/stackups/dart-vs-javascript)  
- [comparing-dart-typescript](https://blog.logrocket.com/comparing-dart-typescript)  
- [dart-vs-typescript](https://aglowiditsolutions.com/blog/dart-vs-typescript) 
## II - Langages Back-End

| Langages | Performance | Facilité d'apprentissage | Coût de développement | Durabilité à long terme | Temps de développement | Consommation d'énergie | Popularité du langage | Syntaxe et lisibilité | Maîtrise du développeur | Résultat |
| -------- | ----------- | ------------------------ | --------------------- | ----------------------- | ---------------------- | ---------------------- | --------------------- | --------------------- | ----------------------- | -------- |
| Java     | 5           | 3                        | 4                     | 5                       | 4                      | 4                      | 5                     | 4                     | 4                       | 38       |
| Go       | 5           | 3                        | 4                     | 4                       | 3                      | 5                      | 4                     | 4                     | 2                       | 34       |
| Python   | 4           | 5                        | 3                     | 5                       | 5                      | 3                      | 5                     | 5                     | 3                       | 38       |
| Rust     | 5           | 2                        | 3                     | 5                       | 3                      | 5                      | 3                     | 3                     | 0                       | 29       |


**Python**

L'avantage principal de Python est sa syntaxe, elle est claire et lisible. Il dispose de frameworks très populaire et robustes comme Django (orienté MVC) et Flask (micro-framework) qui s’adaptent aux petites applications jusqu'au grandes structures complexes. Le problème de Python c'est sa performance sur des applications multi-thread en raison de son GIL (Global Interpreter Lock). Il a cependant un écosystème très développé et une large communauté.

**Java**

Java est un langage orienté objet robuste, il est apprécié pour sa stabilité et surtout pour sa sécurité. Il est en général privilégié dans les grandes entreprises. Les frameworks comme Spring permettent de construire des applications complexes et évolutives, ce qui fait de Java un choix privilégié pour des projets nécessitant une gestion de transactions lourdes. Cependant, Java peut être verbeux contrairement aux langages et est plus complexe à appréhender, bien qu’il bénéficie d’une communauté massive et de nombreux outils.

**Go**

Go (ou Golang) se distingue par sa simplicité et sa performance, il est parfait pour les systèmes distribués et les microservices. Avec son support natif de la concurrence (goroutines), Go est idéal pour les applications qui nécessitent une grande scalabilité. Il a cependant un écosystème plus jeune contrairement aux autre langages de cette liste, certains outils et frameworks ne sont pas encore aussi matures que ceux des autres langages.

**Rust**

Rust a été créé pour résoudre des problèmes complexes de gestion de mémoire et de performance, il est devenu l’un des langages les plus appréciés dans le monde du développement logiciel, avec un statut "langage le plus aimé" dans le sondage Stack Overflow 2023. Rust propose une syntaxe moderne, une sécurité mémoire stricte, et des capacités avancées de concurrence, offrant ainsi des performances robustes pour le développement de systèmes et le back-end. Son adoption par des entreprises majeures comme Amazon, Discord, et Cloudflare témoigne de sa puissance dans des environnements exigeants. Il se démarque par son modèle de propriété, qui garantit la sécurité de la mémoire sans avoir recours à un ramasse-miettes (garbage collector). Ce modèle, combiné avec un typage statique, permet de détecter les erreurs avant la compilation, minimisant les risques d’erreurs mémoire et améliorant la fiabilité des applications. Il offre des abstractions zéro-coût, optimisant les performances sans sacrifier la lisibilité du code. Ses fonctionnalités modernes lui permettent de réduire le temps d'exécution des applications, ce qui le rend adapté aux projets nécessitant des performances élevées. Il a une communauté active, avec une documentation complète, des forums de soutien.

**Source :**

- <a href="https://roadmap.sh/backend/languages" target="_blank">Roadmap.sh</a>
- [choosing-the-right-backend-framework](https://devtoys.io/2024/06/23/choosing-the-right-backend-framework-a-comparison-of-c-python-node-js-go-rust-c-and-java)  
- [backend-development-showdown](https://medium.com/@neeraj_swe/backend-development-showdown-node-js-vs-django-vs-spring-boot-vs-golang-vs-ruby-on-rails-609fcf351f6a)  
- [comparing-backend-architectures](https://inducedofficial.medium.com/comparing-backend-architectures-in-spring-boot-node-js-python-django-and-golang-b4bb0a09bee1) 

## III - Frameworks Front-End

| Framework     | Performance | Écosystème | Développement | Architecture | Déploiement | Sécurité | Coûts | Mobile/Responsive | Résultat |
| ------------- | ----------- | ---------- | ------------- | ------------ | ----------- | -------- | ----- | ----------------- | -------- |
| Vue.js        | 4           | 4          | 5             | 4            | 5           | 4        | 5     | 4                 | 35       |
| Svelte        | 5           | 3          | 4             | 3            | 4           | 3        | 5     | 4                 | 31       |
| Angular       | 4           | 5          | 3             | 5            | 5           | 5        | 4     | 5                 | 36       |
| Flutter (Web) | 3           | 3          | 4             | 4            | 4           | 4        | 5     | 5                 | 32       |


**Angular**

C'est un framework Front-End complet, il est souvent privilégié pour des applications de grande envergure qui nécessitent un support robuste et une architecture complexe. Développé par Google, Angular utilise TypeScript, comme nous l'avons souligné précédemment ça favorise la détection des erreurs lors du développement. Son architecture est modulaire et permet de gérer des fonctionnalités riches et d'assurer une évolutivité. Il a cependant une courbe d'apprentissage plutôt raide, ce qui peut nécessiter un certain temps d'adaptation.

**Vue**

est un framework progressif, créé par <a href="https://evanyou.me/" target="_blank">Evan You</a>, qui allie les avantages d’Angular et la librairie React. Léger et rapide à télécharger, Vue est facile à prendre en main, ce qui en fait un excellent choix pour les développeurs débutants ou pour des projets nécessitant une application légère et performante. Vue utilise un DOM virtuel et prend en charge la liaison de données bidirectionnelle, ce qui le rend efficace pour gérer des applications de taille moyenne. Cependant, bien que sa communauté soit en pleine croissance, elle reste plus petite que celle d’Angular et de React, ce qui limite les ressources d’apprentissage pour les grands projets.

**Svelte**

Svelte propose une approche radicalement différente des autres frameworks en déplaçant la majorité du travail vers une étape de compilation. Contrairement aux autres frameworks qui utilisent le DOM virtuel, Svelte compile directement en JavaScript pur, ce qui optimise les performances et réduit le poids de l’application. Son code est naturellement réactif sans besoin de syntaxe spéciale, et son apprentissage est facilité par l'utilisation de HTML, CSS, et JavaScript/TypeScript standards. Cependant, Svelte a une communauté plus petite et un écosystème limité, ce qui peut poser des défis pour les projets plus complexes nécessitant des bibliothèques tierces.

**Sources :**
- [frontend-languages-ultimate-directory](https://proxify.io/articles/frontend-languages-ultimate-directory)  
- [alternatives-to-javascript](https://cleveroad.com/blog/check-out-what-alternatives-to-javascript-can-be-found-in-the-market)  

## IV - Frameworks Back-End

| Framework          | Performance | Écosystème | Développement | Architecture | Déploiement | Sécurité | Coûts | Mobile/Responsive | Résultat |
| ------------------ | ----------- | ---------- | ------------- | ------------ | ----------- | -------- | ----- | ----------------- | -------- |
| Django (Python)    | 3           | 5          | 4             | 5            | 5           | 5        | 5     | 4                 | 36       |
| Express (Node.js)  | 4           | 5          | 5             | 4            | 5           | 4        | 5     | 4                 | 36       |
| Spring Boot (Java) | 5           | 4          | 3             | 5            | 4           | 5        | 4     | 4                 | 34       |
| Gin (Go)           | 5           | 3          | 3             | 4            | 5           | 5        | 4     | 4                 | 33       |


**Django**

Django est le framework qui correspond aux applications qui nécessitent une grande sécurité et une scalabilité robuste. Il a une approche "batteries-included", il fournit plusieurs outils pour gérer différents aspects d'une application : l'authentification, la gestion de bases de données et le SEO par exemple. Il est notamment apprécié dans les secteurs qui nécessitent des applications rapides à mettre en œuvre, comme les médias sociaux ou les plateformes de contenu.

**Express.js**

Express est le framework de l'environnement NodeJS (Javascript). C'est un framework minimaliste, il très performant et flexible et souvent choisi pour des applications avec la pile de technologies MEAN (MongoDB, Express.js, AngularJS et Node.js) et MERN (MongoDB, Express, React et Node.js). Express est un framework simple d'utilisation. Sa simplicité et sa compatibilité avec JavaScript côté serveur et client facilitent le développement full-stack. Express est idéal pour les API REST et les applications qui nécessitent du temps réel.

**Spring Boot**

Spring Boot est le framework de l'environnement Java, c'est un framework open-source puissant et généralement pour les applications d'entreprise. Il dispose d'une gestion flexible des dépendances et un support étendu pour les microservices, il est couramment utilisé pour les projets nécessitant une scalabilité horizontale (la scalabilité horizontale signifie qu'une application peut gérer plus de trafic ou de charge en ajoutant davantage de serveurs, plutôt qu'en augmentant la puissance d'un seul serveur) et une forte sécurité.

**Gin**

Go est un langage qui ne nécessite pas de framework, son environnement et ses fonctionnalités lui permettent de réaliser des applications sans passer par un framework. Cependant pour les besoin de notre analyse nous présentons tout de même l'un de ces frameworks les plus connus : Gin.

Gin est un framework minimaliste orienté vers la rapidité et l'efficacité, il parfait pour les microservices et les applications nécessitant un support de traitement rapide. Son architecture est légère et son système de middleware le rendent compétitif pour les API nécessitant des performances optimales.

**Sources :**
- [backend-frameworks-list](https://daily.dev/blog/backend-frameworks-list-choosing-the-right-one)  
- [choosing-the-right-backend-framework](https://devtoys.io/2024/06/23/choosing-the-right-backend-framework-a-comparison-of-c-python-node-js-go-rust-c-and-java)  
- [backend-development-showdown](https://medium.com/@neeraj_swe/backend-development-showdown-node-js-vs-django-vs-spring-boot-vs-golang-vs-ruby-on-rails-609fcf351f6a)
## V - Bases de données

| Base de Données    | Performance | Scalabilité | Fiabilité | Modèle de Données | Sécurité | Coûts | Administration | Intégration | Support et Communauté | Résultat |
| ------------------ | ----------- | ----------- | --------- | ----------------- | -------- | ----- | -------------- | ----------- | --------------------- | -------- |
| PostgreSQL         | 5           | 4           | 5         | 4                 | 5        | 4     | 4              | 5           | 5                     | 41       |
| MongoDB            | 4           | 5           | 4         | 5                 | 4        | 4     | 4              | 5           | 5                     | 40       |
| MySQL              | 4           | 4           | 5         | 3                 | 5        | 5     | 5              | 4           | 5                     | 40       |
| Firebase Firestore | 3           | 5           | 4         | 4                 | 5        | 3     | 4              | 5           | 4                     | 37       |
| DynamoDB           | 4           | 5           | 5         | 4                 | 5        | 3     | 3              | 5           | 4                     | 38       |


**Les base de données SQL & NoSQL**

Les bases de données SQL et NoSQL diffèrent principalement dans leur modèle de données, leur flexibilité, et leurs cas d’utilisation privilégiés. Les bases de données SQL, aussi appelées relationnelles, sont structurées en tables avec des lignes et des colonnes, permettant de définir des relations précises entre les données via des clés étrangères et des jointures. Ce modèle favorise une organisation rigide qui garantit l’intégrité et la consistance des données grâce au respect des propriétés ACID (Atomicité, Cohérence, Isolation, Durabilité). SQL est ainsi idéal pour des applications où la précision des transactions est cruciale, comme les systèmes financiers et les e-commerces.

Les bases de données NoSQL, également appelées non relationnelles, adoptent une approche plus flexible en stockant les données sous forme de documents, de paires clé-valeur, de colonnes, ou de graphes. Cette structure est adaptée aux données non structurées ou semi-structurées, ce qui permet d’évoluer et de s’adapter facilement aux modifications de structure. De plus, les bases de données NoSQL sont optimisées pour l’évolutivité horizontale (cf: explications du framework Srping Boot), ce qui signifie qu’elles peuvent gérer des volumes de données massifs en répartissant la charge sur plusieurs serveurs. Ces bases de données sont particulièrement adaptées aux applications nécessitant une haute disponibilité et une flexibilité de stockage, comme les réseaux sociaux, le big data, et les applications IoT.

**MongoDB**

MongoDB est une base de données NoSQL, idéale pour les données non structurées et les applications nécessitant une grande flexibilité et une évolutivité horizontale. Elle stocke les données sous forme de documents JSON, ce qui facilite l’ajout ou la modification de champs. MongoDB est souvent privilégiée pour les applications de gestion de contenu, d'e-commerce ou d'analyses en temps réel. Elle offre une haute disponibilité grâce à la réplication et une mise à l'échelle aisée via le partitionnement, ce qui en fait un bon choix pour les projets à fort volume de données.

**PostgreSQL**

PostgreSQL est une base de données relationnelle robuste, conforme aux normes ACID, qui favorise l'intégrité des données et supporte des requêtes SQL complexes. Conçue pour des données structurées, elle est particulièrement efficace pour les applications transactionnelles et d'e-commerce, grâce à des fonctionnalités avancées de gestion des transactions et d'indexation. PostgreSQL est aussi extensible et supporte JSON, offrant une certaine flexibilité pour des cas d'usage hybrides avec des données semi-structurées.

**MySQL**

MySQL est également une base de données relationnelle, elle est largement utilisée pour sa simplicité et sa fiabilité. Elle convient bien aux applications web de petite à moyenne envergure, comme les systèmes de gestion de contenu ou les applications de type blog. MySQL permet une mise à l'échelle horizontale via le partitionnement et la réplication, mais peut être limitée pour des charges de travail plus complexes nécessitant des jointures et des transactions ACID complexes.

**Firebase**

Firebase est une base de données NoSQL qui se distingue par ses capacités de synchronisation en temps réel et son intégration avec des applications mobiles, notamment via la Google Cloud Platform. Elle est idéale pour les jeux, le chat en direct et les réseaux sociaux, avec des fonctionnalités comme le mode hors ligne et une gestion simplifiée des utilisateurs et des notifications. Cependant, Firebase est fortement liée à l’écosystème Google, ce qui peut poser des défis en matière de migration vers d'autres plateformes.

**DynamoDB**

DynamoDB une autre base de données NoSQL d'Amazon Web Services (AWS), est conçue pour des applications nécessitant une évolutivité et des performances élevées, comme l'IoT ou les analyses de flux de données. Elle supporte un modèle de clé-valeur flexible et propose une scalabilité automatique en fonction de la charge de travail. DynamoDB est adaptée aux charges de travail imprévisibles et s'intègre parfaitement aux autres services AWS, bien qu’elle soit limitée pour les requêtes SQL complexes.

**Sources :**
- [backend-frameworks-list](https://daily.dev/blog/backend-frameworks-list-choosing-the-right-one)  
- [choosing-the-right-backend-framework](https://devtoys.io/2024/06/23/choosing-the-right-backend-framework-a-comparison-of-c-python-node-js-go-rust-c-and-java) 
- 
## VI - Hébergements Cloud Front-End

| Service        | Performance | Facilité d'utilisation | Coût | Scalabilité | Sécurité | Temps de déploiement | Popularité | Intégration CI/CD | Résultat |
|--------------|------------|----------------------|------|-------------|------------|------------------|------------|----------------|----------|
| Vercel      | 5          | 5                    | 4    | 5           | 5          | 5                | 5          | 5              | 39       |
| Netlify     | 4          | 5                    | 4    | 4           | 5          | 5                | 5          | 5              | 37       |
| Firebase    | 4          | 4                    | 3    | 4           | 5          | 5                | 5          | 4              | 34       |
| Cloudflare Pages | 4     | 5                    | 5    | 5           | 5          | 5                | 4          | 4              | 37       |
| AWS Amplify | 5          | 3                    | 3    | 5           | 5          | 4                | 4          | 5              | 34       |

**Vercel**

Vercel est une solution d’hébergement optimisée pour les applications front-end, notamment celles basées sur Next.js. Il offre un déploiement simple, une mise en cache performante et une excellente intégration avec GitHub.

**Netlify**

Netlify est très populaire pour le déploiement rapide d’applications statiques et de sites Jamstack. Il propose un système de build automatisé et une intégration facile avec des CMS headless.

**Firebase**

Firebase Hosting est conçu pour héberger des applications web avec des fonctionnalités en temps réel, idéal pour les applications avec une forte interaction utilisateur et des bases de données NoSQL.

**Cloudflare Pages**

Cloudflare Pages est un service d’hébergement performant qui exploite le réseau Cloudflare pour proposer un chargement rapide et une haute sécurité.

**AWS Amplify**

AWS Amplify est conçu pour le développement et le déploiement d’applications front-end et mobiles, avec une forte intégration aux services AWS comme Cognito et DynamoDB.

**Sources :**

- [Comparaison des services d'hébergement cloud front-end : Vercel, Netlify, Firebase, Cloudflare Pages, AWS Amplify](https://david.dev/netlify-vs-vercel-vs-coudflare-pages-vs-google-cloud-vs-aws-amplify)

## VII - Hébergements Cloud Back-End

| Service          | Performance | Facilité d'utilisation | Coût | Scalabilité | Sécurité | Temps de déploiement | Popularité | Intégration CI/CD | Résultat |
| ---------------- | ----------- | ---------------------- | ---- | ----------- | -------- | -------------------- | ---------- | ----------------- | -------- |
| Railway          | 4           | 5                      | 4    | 5           | 5        | 5                    | 4          | 5                 | 37       |
| Render           | 4           | 5                      | 4    | 5           | 5        | 5                    | 4          | 5                 | 37       |
| Fly.io           | 5           | 4                      | 4    | 5           | 5        | 4                    | 4          | 5                 | 36       |
| AWS Lambda       | 5           | 3                      | 3    | 5           | 5        | 4                    | 4          | 5                 | 34       |
| Google Cloud Run | 5           | 4                      | 3    | 5           | 5        | 4                    | 4          | 5                 | 35       |

**Railway**

Railway permet un déploiement rapide et sans configuration, idéal pour les startups et les développeurs qui souhaitent déployer rapidement des applications back-end sans se soucier de l'infrastructure.

**Render**

Render est une plateforme similaire à Railway qui propose des déploiements simples avec un excellent rapport performance-coût, adapté aux applications web et API.

**Fly.io**

Fly.io est conçu pour héberger des applications proches des utilisateurs avec une architecture distribuée, parfait pour des applications serverless et edge computing.

**AWS Lambda**

AWS Lambda est une solution serverless qui permet d'exécuter du code en réponse à des événements sans gérer d’infrastructure, optimisé pour des applications event-driven.

**Google Cloud Run**

Google Cloud Run est un service serverless permettant d’exécuter des conteneurs en réponse à des requêtes HTTP, adapté aux applications modernes et microservices.

**Sources :**

- [Comparaison des services d'hébergement cloud back-end : Railway, Render, Fly.io, AWS Lambda, Google Cloud Run](https://ritza.co/articles/gen-articles/render-vs-heroku-vs-vercel-vs-railway-vs-fly-io-vs-aws/)

## VIII - Choix final

### Langage Front-End

![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)

Le langage **TypeScript** a été choisi pour le développement Front-End en raison de sa robustesse et de sa capacité à améliorer la qualité du code. TypeScript est une extension de JavaScript qui introduit le typage statique, permettant de détecter les erreurs dès la compilation et d’assurer une meilleure maintenance à long terme. 

Son adoption permet d’améliorer la productivité des développeurs, en réduisant les erreurs et en facilitant le refactoring du code. Il est compatible avec tous les frameworks modernes, notamment **Svelte**, garantissant une intégration fluide et performante.

<u>Justifications :</u>

- **Sécurité** : Typage statique permettant d’éviter les erreurs courantes.
- **Maintenance** : Code plus lisible et plus structuré, facilitant les évolutions futures.
- **Scalabilité** : Adapté aux projets de grande envergure.

### Framework Front-End

![Svelte](https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white)

Nous avons retenu **SvelteJS** comme framework Front-End pour sa légèreté et son approche innovante. Contrairement aux autres frameworks basés sur un DOM virtuel, Svelte compile le code en JavaScript pur au moment du build, améliorant ainsi les performances d’exécution.

Sa simplicité et son absence de surcharge permettent un développement rapide et efficace, idéal pour un **MVP**. De plus, il favorise la création de **Progressive Web Apps (PWA)**, ouvrant la voie à une transition vers des applications mobiles sans dépendre immédiatement d’un framework mobile dédié.

<u>Justifications :</u>

- **Développement rapide** : Moins de complexité, courbe d’apprentissage réduite.
- **Performances** : Code compilé optimisé, sans surcharge inutile.
- **Flexibilité** : Possibilité de générer des **PWA** pour un déploiement mobile simplifié.

### Langage Back-End

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Le choix du langage **Back-End** s'est porté sur **Golang**, reconnu pour ses **performances élevées**, sa **simplicité** et sa **scalabilité**. Développé par Google, Go est conçu pour offrir une exécution rapide et efficace, tout en minimisant la consommation de ressources.

Sa gestion native de la concurrence grâce aux **goroutines** en fait une option privilégiée pour des applications nécessitant **une forte réactivité**, comme un projet e-sport. Il est particulièrement adapté aux **architectures microservices** et aux **applications distribuées**.

<u>Justifications :</u>

- **Idéal pour les prototypes** : Développement rapide sans surcharge inutile.
- **Performance et scalabilité** : Gestion concurrentielle efficace, idéal pour les applications temps réel.
- **Interopérabilité cloud** : Facile à conteneuriser avec **Docker** et déployable sur **Kubernetes, AWS, GCP, OVHcloud**.
- **Sécurité et robustesse** : Gestion stricte des erreurs, typage fort et exécution compilée.

### Framework Back-End

![Gin](https://img.shields.io/badge/gin-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Le choix du framework **Gin** s’impose naturellement en complément de **Golang**. Gin est un framework HTTP rapide et léger, idéal pour la création d’API performantes. Son architecture repose sur une gestion optimisée des requêtes HTTP, garantissant un traitement rapide des données.

Avec une prise en charge native des **middlewares**, un **routage performant**, et une **gestion des erreurs simplifiée**, Gin permet de construire une **API REST robuste** en un temps réduit, tout en assurant une maintenance facile.

<u>Justifications :</u>

- **Ultra-performant** : Faible latence et haute capacité de traitement.
- **Minimaliste mais extensible** : Offre un socle léger avec des fonctionnalités modulables.
- **Adapté aux microservices** : Intégration fluide avec des solutions cloud-native.

### Base de données

![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

La base de données choisie est **MongoDB**, reconnue pour sa flexibilité et son adaptabilité aux **données non structurées**. Grâce à son modèle orienté **documents (JSON-like)**, elle permet une gestion fluide des données dynamiques, idéale pour un projet e-sport où les structures de données peuvent évoluer.

MongoDB facilite également la **scalabilité horizontale**, assurant une gestion efficace des montées en charge lorsque l’application grandit. Elle est optimisée pour le **stockage des objets complexes**, comme les **profils de joueurs**, les **statistiques de performance**, ou les **roadmaps personnalisées**.

<u>Justifications :</u>

- **Flexibilité** : Modèle documenté dynamique, adapté aux applications évolutives.
- **Simplicité** : Facilité d’utilisation et d’intégration avec **Golang**.
- **Performance** : Optimisé pour le traitement rapide de requêtes complexes.
- **Scalabilité horizontale** : Peut gérer un grand volume de données sans perte de performance.

### Hébergement Cloud Front-End

![Vercel](https://img.shields.io/badge/vercel-%23000000.svg?style=for-the-badge&logo=vercel&logoColor=white)

Pour l’hébergement du **Front-End**, **Vercel** est le choix privilégié. Conçu pour offrir une **expérience fluide aux développeurs**, Vercel permet un déploiement en **quelques secondes** avec un support natif des frameworks modernes, notamment **Svelte**.

Son infrastructure optimisée et sa mise en cache avancée garantissent une **latence minimale**, idéale pour un projet nécessitant **des temps de réponse rapides**.

<u>Justifications :</u>

- **Déploiement rapide et automatisé** : Intégration directe avec Git.
- **Performance** : Optimisé pour les applications web modernes.
- **Scalabilité automatique** : Gère les pics de trafic sans intervention manuelle.

### Hébergement Cloud Back-End

![Railway](https://img.shields.io/badge/railway-%23005990.svg?style=for-the-badge&logo=railway&logoColor=white)

L’hébergement du **Back-End** repose sur **Railway**, une plateforme qui permet de **déployer rapidement des applications serveur** sans gestion complexe d’infrastructure. Railway offre une **expérience simplifiée** et une **scalabilité immédiate**, adaptée aux besoins d’un projet en phase MVP.

Avec un support natif de **Docker** et une intégration fluide avec **MongoDB**, Railway garantit une mise en production rapide, sécurisée et optimisée.

<u>Justifications :</u>

- **Déploiement sans effort** : Configuration minimale, mise en production immédiate.
- **Adapté aux microservices** : Idéal pour une architecture modulaire.
- **Optimisé pour Golang et MongoDB** : Support natif des bases de données et des conteneurs.