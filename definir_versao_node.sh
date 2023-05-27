#!/bin/bash

# Verificar se o arquivo package.json existe no diretório atual
if [ -f "package.json" ]; then
  # Ler a versão do Node.js especificada no arquivo package.json
  version=$(node -pe "require('./package.json').engines.node")

  # Verificar se a versão foi encontrada
  if [ -n "$version" ]; then
    echo "Versão do Node.js especificada no package.json: $version"

    # Definir a versão do Node.js usando o npm use
    npm use $version

    echo "Versão do Node.js definida com sucesso!"
  else
    echo "Versão do Node.js não encontrada no package.json"
  fi
else
  echo "Arquivo package.json não encontrado no diretório atual"
fi
