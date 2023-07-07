# go-fiber-api



-> Explicar o que é o CONTEXTO
Você utiliza ele quando quer utilizar algo de fora, por exemplo ->
* Limitar o tempo de chamada, aahhh quero que o httpget feito demore no máximo 1s, se não cancele.
* Ahh valide o campo xyz de outra api quando eu chamar a principal
* Ahh quero que quando eu chamar a api principal, ele chame outras apis e retorne o resultado de todas elas

Define um tipo de contexto -> Deadlines, sinais de cancelamento e outros valores com escopo de solicitação
entre os limites da API e entre os processos.

Utilizam channels por debaixo, eles ficam "ouvindo os sinais que recebem do ctx"

3 tipos de contexto - >
Deadline: Tempo MÁXIMO, se passar, cancela
Cancellation Signals: Código, cancela a execução caso role algo
Request-scoped values: Manda um chave valor no handler por exemplo pra uma goroutine, ai, o valor definido
é assincrono e imutável.

WITHDEADLINE -> é um dos metodos de CONTEXT -> você passa o horário MÁXIMO
-> recebe -> parent context.Context, t time.time
-> retorna -> context.Context, cancel func()

WITHTIMEOUT -> é um dos metodos de CONTEXT. -> você passa o tempo MÁXIMO, DURAÇÃO
-> recebe -> parent context.Context, timeout time.Duration
-> retorna -> context.Context, cancel func()



CONSTRUTOR PARA ACESSO OU VARIÁVEL GLOBAL
-> construtor - newmongodbconnection (blablabla) return mongo client
utilizar o próximo

var (connection mongo.client)

MARHSAL
-> TEM COMO FUNÇÃO -> transformar um objeto em um array de bytes, UMA STRING EM BYTES E QUALQUER OUTRA COISA
no caso, utilizamos comummente para trasnformar um struct em bytes.

UNMARSHAL
-> TEM COMO FUNÇÃO -> transformar um array de bytes em um objeto, UM ARRAY EM OBJETO OU ETC


-> aula 13 - https://www.youtube.com/watch?v=wbtdg9cR3pY&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr&index=14

https://www.youtube.com/watch?v=ywLyAvA9e_8&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr&index=6