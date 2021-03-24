Accounts

# Swarm

### export env variables
```bash
export $(cat .env.prod)
```
### deploy
```bash
 docker stack deploy -c stack.yml accounts
```
