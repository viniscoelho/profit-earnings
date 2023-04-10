### Descrição da resolução
De modo geral, não utilizei nenhuma biblioteca externa (a não ser a testify para os testes unitários).
Fiz algumas customizações, como a criação de um tipo `RoundedFloat64` para facilitar a geração do output.
Basicamente, quando eu faço um marshaling da estrutura `Tax`, eu não preciso tratar a precisão de ponto flutuante.
Caso contrário, eu teria que fazer um for loop e fazer isso no print manualmente. Não que ficasse ruim a leitura,
mas preferi dessa forma para ficar mais simples (principalmente se precisar estender a saída para uma API).
Além disso, organizei o código em pequenos pacotes para facilitar a separação de responsabilidades. Por exemplo,
o `main.go` fica na raíz e na pasta `src`, contém o código divido em `operations`(onde a estrutura `Earnings` criada a cada novo conjunto de operações) e `types`, onde fica os tipos comuns usados no `main.go` e `calculate_earnings.go`.

Como descrito na especificação, considerei que os casos de entrada não quebrariam no momento de leitura, então,
utilizei os exemplos oferecidos para testar.

### Como usar
- Uma das formas de executar é compilando o projeto com:
    - `go build -mod vendor -o profit-earnings`
    - em seguida, executar o script `run_test_cases.sh`
        - será necessário dar permissões de execução ao script (chmod u+x run_test_cases.sh)
        - ou executar com `bash run_test_cases.sh`
    - Contudo, para a execução desses passos, é necessário instalar as dependências de Golang. 

- A outra forma de através do uso de docker:
  - `docker build -t profit-earnings .` e então,
  - `docker run profit-earnings`

Dessa forma, o dockerfile vai criar todas as dependências necessárias e copiar os arquivos necessários para a execução.
Os casos de teste enviados na descrição do desafio foram copiados para a pasta `test_cases` e são executados pelo script `run_test_cases.sh`. Dessa forma, não é necessário rodar os testes um a um, nem mesmo instalar Golang na máquina (caso não tenha).

Os testes unitários são os casos de teste comparados com as saídas esperadas, as mesmas enviadas no arquivo contendo a descrição do desafio. Para executar, no entando, é necessário que tenha o Golang instalado.
Para isso, navegue até a pasta raiz do projeto e execute no terminal:
`go test ./src/operations/`