

az containerapp up \
  --name htmx-go-todo \
  --resource-group temp \
  --location centralus \
  --environment 'my-aca' \
  --image ghcr.io/benc-uk/htmx-go-todo:latest \
  --target-port 4000 \
  --ingress external \