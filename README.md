# etcd-playground 🔏
An ETCD playground to learn it more in depth

## Usage 🖊

1) Start the cluster
```bash
docker compose up -d
```

2) Export endpoints to a gloabl var to be used in the CLI commands
```bash
export ETCD_ENDPOINTS="localhost:32771,localhost:32772,localhost:32773"
```

3) Put KV
```bash
etcdctl --endpoints="$ETCD_ENDPOINTS" put foo bar --prev-kv --write-out=json | jq
```

4) Get KV
```bash
etcdctl --endpoints="$ETCD_ENDPOINTS" get --prefix foo --write-out=json | jq
```

5) Watch KV
```bash
etcdctl --endpoints="$ETCD_ENDPOINTS" watch foo --prev-kv --write-out=json | jq
```
