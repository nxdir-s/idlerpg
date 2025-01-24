yq -i '(select(.data[]) | .data[]) style="double"' ./dist/0001-idlerpg.k8s.yaml
