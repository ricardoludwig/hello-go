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

## Go Build

- gerar o executável
- gerar o executável para uma plataforma específica (cross compilation)
	GOOS=windows GOARCH=386 go build -o hello.exe

## Arrays

	var frutas [3] string
	frutas[0] = laranja
	frutas[2] = banana
	frutas[3] = maçã

	len(frutas)

	for index, fruta := range frutas {
		fmt.Printf("Fruta[%d] é %s/n/r", index, fruta)
	}

## Slices

- pode ser entendido como um array dinâmico, semelhante ao List em Java

	var nums []int
	fmt.Println(nums, len(mums), cap(nums))
	//out -> [] 0 0

	//inicializando o slice
	nums = make([]int, 5)
	fmt.Println(nums, len(mums), cap(nums))
	//out -> [0 0 0 0 0] 5 5

	nums = append(nums, 6)

- pegando dados em intervalos do slice (dividindo o slice)

	// Primeiro item começa com o indice 0
	// Segundo item começa com o indice 1
	temp1 := nums[2:5]

	//Dividindo a partir do primeiro item (dois primeiros itens)
	temp2 := nums[:2]

	//Pegando os dois últimos itens
	temp3 := nums[len(nums) - 2:]

- removendo item do slice

	É preciso obter o indice do item que deseja remover, após isso criar um Slice temporário sem o item que deseja-se remover e por fim
	copiar o Slice temporário para o Slice original

	indexItemRemover := 2
	temp := cidades[:indexItemRemover]
	temp = append(temp, cidades[indexItemRemover+1:]...)
	copy(cidades, temp)

## Interfaces

- interfaces são um componento no Go que nos ajuda a escrever códigos mais complexos, quando desejamos reaproveitar códigos.
- no Go a herança não é direta, existe uma herança por semelhança, isso se dá pela implementação das ações que te caracterizam
como pertencente aquela interface. Se em uma dada Struct implementarmos funções de uma dada interface o compilador irá inferir que estamos
implementando dada interface, não há a necessidade do uso de palavras chaves como implements, por exemplo.
- Go não é uma linguagem 100% Orientada à Objetos.

## Comando defer

- pode ser declarado logo depois da abertura de algum recurso (escrita de arquivo, conexão com banco) que após a execução da função o
comando defer será executado para liberar o recurso. Isso é muito útil para evitar esquecimentos dos programadores em fechar recursos.




































