# Mods API (WIP)
Simple API hosted on GAE with the intention to easily extend it with more
modules.

## TODO
* Wrap the Request and give utility methods (e.g. ParseBody(*interface))
* Make a Server "class"
* Define Http Method in endpoints and handle `Method Not Allowed` globally
* More modules
* Code origanisation (have modules in folders to split their code?)
* Maybe: use a more elaborate router (gorilla/mux?)