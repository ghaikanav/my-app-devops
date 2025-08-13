ENV_VALUES_FILE="values.yaml"
helm upgrade --install my-app . -f "$ENV_VALUES_FILE"