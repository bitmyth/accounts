Accounts

# Deploy to Swarm

Add label role=db to node where mysql will runs
```bash
docker node update --label-add role=db NODE
```

### export env variables
```bash
export $(cat .env.prod)
```
### deploy
```bash
 docker stack deploy -c stack.yml accounts
```
