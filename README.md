# TP Go - Programmation répartie

### Build le projet

```
.\gobuild.bat
```
Génère un fichier restserv.exe

### Lancer le projet

```
.\restserv.exe
```

### Entité

__Language :__  

- Code
- Name

__Student :__  

- Id
- FirstName
- LastName
- Age
- LanguageCode

### Utilisation de l'API REST

Le serveur est lancé en local sur le port 9999.

_{entity}_ est à remplacer sois par _students_ sois par _languages_  
_{key}_ est à remplacer sois par l'_Id_ du student sois par le _Code_ du language  


- {GET} /{entity} : Récupère toutes les entités 

- {GET} /{entity}/{key} : Récupère l'entité correspondant à la clé

- {DELETE} /{entity}/{key} : Supprime l'entité correspondant à la clé (un _Code_ ou un _Id_)

- {POST} /{entity} : Créer l'entité via un JSON passé dans le body

- {PUT} /{entity} : Met à jour l'entité via un JSON passé dans le body
