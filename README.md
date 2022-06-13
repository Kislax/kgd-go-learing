# golang-learning

[Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/) exercises repository template.

# About me

Short bio and motivation in learning golang.

# Learned lessons

### arrays - 100.0%

<details>
  <summary><code>func Sum(numbers []int) int</code></summary>

</details>

<details>
  <summary><code>func SumAll(numbersToSum ...[]int) []int</code></summary>

</details>

<details>
  <summary><code>func SumAllTails(numbersToSum ...[]int) []int</code></summary>

</details>

### helloworld - 87.5%

<details>
  <summary><code>func Hello(name string, language string) string</code></summary>

</details>

### integers - 100.0%

<details>
  <summary><code>func Add(x, y int) int</code></summary>

</details>

### hellogo - 100.0%
Package hellogo contains first steps in language.
<details>
  <summary><code>func Hello() string</code></summary>

    Hello is first function.
</details>

### pointers - 85.7%
VARIABLES

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

TYPES

type Bitcoin int

func (b Bitcoin) String() string

type Stringer interface {
	String() string
}

type Wallet struct {
	// Has unexported fields.
}

func (w *Wallet) Balance() Bitcoin

func (w *Wallet) Deposit(amount Bitcoin)

func (w *Wallet) Withdraw(amount Bitcoin) error
### mapsl - 90.0%
VARIABLES

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

TYPES

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error

func (d Dictionary) Search(word string) (string, error)