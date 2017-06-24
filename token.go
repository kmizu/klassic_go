package parser

type TokenType int

type Tokenizer struct {
  input string
  position int
  line int
  column int
}

const (
  TOKEN_INT TokenType = iota
  TOKEN_FLAOT
  TOKEN_IDENT
  TOKEN_IF
  TOKEN_LPAREN
  TOKEN_RPAREN
  TOKEN_LBRACE
  TOKEN_RBRACE
  TOKEN_LBRACKET
  TOKEN_RBRACKET
  TOKEN_WHITE_SPACES
)

func NewTokenizer(input string) *Tokenizer {
  return &Tokenizer{input, 0, 1, 1}
}

func (self Tokenizer) NextToken() {
}
