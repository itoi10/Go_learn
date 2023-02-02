package lexer

import "token_analysis/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の文字を指し示す位置
	readPosition int  // これから読み込む現在の文字の次の位置
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 次の1文字を読むヘルパー関数
func (l *Lexer) readChar() {
	// 終端に達した場合は0(ASCIIのNUL文字)とする
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// 現在位置の文字読み込み
		l.ch = l.input[l.readPosition]
	}
	// ポインタを進める
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// 文字列?
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
			// 数値?
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
			// 未知の語
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// ポインタを次の文字に進めてからトークン返却
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 識別子を読む. 非英字に到達するまで読み進める
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]

}

// 引数が英字か判定する. 英字は_を含む
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '-'
}

// 空白文字を読み飛ばす
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 数値を読む. 非数値に到達するまで読み進める
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 引数が数値か判定する
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
