# Mods API (WIP)
Simple API hosted on GAE with the intention to easily extend it with more
modules.

Hosted at: https://mods-api.appspot.com/

# Examples
```
curl https://mods-api.appspot.com/coin/flip
curl https://mods-api.appspot.com/dice/roll
curl -X POST https://mods-api.appspot.com/decider -d '["apple", "oranges"]'
```

## TODO
* Make a Server "class"
* Define Http Method in endpoints and handle `Method Not Allowed` globally
* More modules
* Code origanisation (have modules in folders to split their code?)
* Maybe: use a more elaborate router (gorilla/mux?)