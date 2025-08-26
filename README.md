# API de Sincronização de Posts

## Descrição
Uma API em Go que consome o serviço público JSONPlaceholder para buscar posts, persiste os dados em um banco de dados local e expõe endpoints para leitura e manipulação desses registros. 

## Tecnologias Principais
- Go (módulos)
- Chi (roteamento HTTP)
- MySQL (conexão inicial)
- godotenv (carregamento de variáveis de ambiente)

## Estrutura de Pastas (resumo)
```
cmd/api          # Ponto de entrada da aplicação
configs          # Arquivos de configuração (.env)
internal/db      # Conexão com banco de dados
internal/domain  # Entidades de domínio
internal/repository # Implementações de repositórios (API externa, memória)
internal/service # Regras de negócio
internal/http    # Handlers e roteador
migrations       # Scripts SQL de migração
```

## Variáveis de Ambiente
Criar um arquivo `configs/.env` baseado em `.env.example`:
```
POSTS_API_URL=https://jsonplaceholder.typicode.com
```

## Executando o Projeto
1. Clonar o repositório:
```
git clone https://github.com/matheuszuge/go-posts-sync.git
cd go-posts-sync
```
2. Criar o arquivo de ambiente:
```
cp configs/.env.example configs/.env
```
3. (Opcional) Ajustar `POSTS_API_URL` se necessário.
4. Garantir que o MySQL esteja acessível conforme DSN definido em `internal/db/db.go`.
5. Rodar a aplicação:
```
go run ./cmd/api
```

A API sobe por padrão em `http://localhost:8080`.

## Endpoints Atuais
| Método | Caminho    | Descrição                  |
|--------|----------- |--------------------------- |
| GET    | /ping      | Health-check simples       |
| GET    | /users     | Lista usuários (in-memory) |
| GET    | /posts     | Lista posts da API externa |

## Migrações
Os arquivos SQL estão em `migrations/`. Para aplicar manualmente, executar os scripts `*.up.sql` na ordem numérica em seu banco de dados. 

## Convenções de Código
- Pacotes dentro de `internal` para encapsulamento.
- Handlers somente com lógica de transporte (serialização, status HTTP).
- Serviços contêm regras de negócio.
- Repositórios isolam detalhes de acesso a dados externos ou internos.

## Licença
MIT
