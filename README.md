Repositório para estudo e prática em Go.

# Anotações

## Variáveis

- fortemente tipdas;
- caracateres são unicode;
- rune é tipo utilizado para caracateres especiais;
- não existe valor nulo para variáveis de tipos primitivos;
- se a variável não tiver o tipo explicitamente declarado o compilador será capaz de inferir seu tipo baseado no valor que for atribuído;
- toda variável declarada deve ser utilizada, caso não seja acarretará em um erro de compilação;
- em um primeiro momento é uma linguagem é uma linguagem procedural, estruturada. Depois novos recursos foram sendo adicionados. Ex. Inter-
farces, funções anônimas, clousures. É uma linguagem focada na simplicidade;


### Escopo

- escopo de pacote,variáveis podem ser públicas ou privadas. Públicas são declaradas com a primeira letra maiúscula e são acessíveis por
outros pacotes (variável global);
- escopo de função, acessíveis apenas não funções que foram declaradas;

[Variáveis](variaveis/variavel.go)

## Pacote

- em go nõa é permitida a dependência cíclica. Ex. A ┴ B mas B não pode depender de A

## Funções

- para exportar uma função (torná-la pública) é necessário iniciar seu nome com a primeira letra maiúscula;
- função é um tipo de dado em Go, ou seja eu posso passá-la como parâmetro;
- podem ter múltiplos retornos;
- função com o retorno assinado, já inicializa a variável que será retornada e melhora a legibilidade do cógido.
Ex. func A(x int) (y int) { ... }

## Struct

- por convenção não precisamos informar o nome dos atributos do Struct na declaração, mas a ordem deve seguir a ordem que os atributos
foram declarados para o compilador possa inferir corretamente as atribuições
- é possível atrelar funcões à uma Struct, exemplo:

	type Endereco struct {
		logradouro string
		numero int
	}

	func (e *Endereco) setLogradouro (log string) {
		e.logradouro = log
	}
- em Structs podemos utilizar tag para identificação de seus atributos, um exemplo e transfomação em um JSON.

	type Endereco struct {
		logradouro `json:"rua"
	}

## Ponteiros

- por padrão a passagem de parâmetros em Go é por cópia
- a utilização dos ponteiro é útil quando desejamos passar uma variável por referência para uma função


## Condicionais

### if

- no exemplo abaixo é demosntrado uma forma otimizada do if

	//variaveis x, status fazem parte apenas do escopo do if
	if x, status := fnRetornaUmBoleano(z); satus {
			fmt.Println("Status true")
			return
	}


## Loop

### Switch
- Go ao contrário de outras linguagens não passa por todas as condições do switch, para esse comportamento é utilizado a palavra chave
fallthrough

## Map

- Coleção de dados com uma chave, o algorítimo do Maps em Go é muito eficiente
- map se não inicializado fica nulo
- não é recomendado utilizar um struct como chave, isso degrada o desempenho, é recomendado a utilização de string ou int

## Tratamento de erro

- erro é um ponteiro de um objeto que implementa a interface Error
- é recomendado criar tipos de erros personalizados
- é possível avaliar o tipo de erro que está sendo retornado
-



















