# Descrição do Projeto
Este projeto é um exemplo de aplicação cliente-servidor em Golang. 
Ele demonstra a comunicação entre um cliente e um servidor.

>Nota: Este projeto é uma demonstração e não segue todas as boas práticas arquiteturais.

# Como Executar o Projeto
1. Clone o repositório:
   ```sh
   git clone https://github.com/caricciy/client-server
   ```

2. Navegue até o diretório do projeto:
   ```sh
   cd client-server
   ```

3. Configurando arquivo .env:
    Primeiramente, copie o arquivo `example.env` para `.env`:
    ```sh
    cp example.env .env
    ```
    Em seguida, edite o arquivo `.env` e configure as variáveis de ambiente conforme necessário.

3. Execute o servidor:

    Para essa etapa, precisamo utilizar dois terminais.

    Abra o primeiro terminal e execute o servidor:
    ```sh
    make run-server
    ```
    Em outro terminal, execute o cliente:

    ```sh
    make run-client
    ```

# Variáveis de Ambiente

O projeto utiliza as seguintes variáveis de ambiente:

- `SERVER_URL` : URL do servidor que executamos com o comando `make run-server` (padrão: `http://localhost:8080`).
- `SERVER_PORT` : Porta do servidor que executamos com o comando `make run-server` (padrão: `8080`).
- `EXT_API_URL` : API de terceiros utilizada para obter os dados de cambio (padrão: `https://economia.awesomeapi.com.br/json/last/USD-BRL`).
- `DB_DIR` : Diretório onde o banco de dados SQLite será salvo.
- `OUTPUT_DIR` : Diretório onde os arquivos de saída gravados pela aplicação cliente que executamos com o comando `make run-client` serão salvos.
