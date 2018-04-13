# Mods API (WIP)
Very Simplistic and dockerzied Golang API

# Examples
```
curl http://localhost:8000/coin/flip
curl http://localhost:8000/dice/roll
curl -X POST http://localhost:8000/decider -d '["apple", "oranges"]'
```

## TODO
* Define Http Method in endpoints and handle `Method Not Allowed` globally
* More modules
