# Mods API (WIP)
Very Simplistic and dockerzied Golang API, currently deployed at http://52.168.124.14/coin/flip.

# Building and running the container
```
make build
make run
```

# Examples
```
curl http://localhost:8000/coin/flip
curl http://localhost:8000/dice/roll
curl -X POST http://localhost:8000/decider -d '["apple", "oranges"]'
```

## TODO
* Define Http Method in endpoints and handle `Method Not Allowed` globally
* More modules
