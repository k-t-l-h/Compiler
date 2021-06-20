// Generated from C.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 40, 233, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 3, 2, 3, 2, 5, 2, 82, 10, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 
	4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 
	3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 
	23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 
	3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 
	27, 3, 27, 3, 28, 3, 28, 7, 28, 168, 10, 28, 12, 28, 14, 28, 171, 11, 28, 
	3, 29, 6, 29, 174, 10, 29, 13, 29, 14, 29, 175, 3, 30, 6, 30, 179, 10, 
	30, 13, 30, 14, 30, 180, 3, 30, 3, 30, 7, 30, 185, 10, 30, 12, 30, 14, 
	30, 188, 11, 30, 3, 30, 3, 30, 6, 30, 192, 10, 30, 13, 30, 14, 30, 193, 
	5, 30, 196, 10, 30, 3, 31, 3, 31, 3, 31, 5, 31, 201, 10, 31, 3, 31, 7, 
	31, 204, 10, 31, 12, 31, 14, 31, 207, 11, 31, 3, 31, 3, 31, 3, 32, 3, 32, 
	3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 
	37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 
	2, 2, 40, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 
	75, 39, 77, 40, 3, 2, 7, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 
	92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 
	15, 15, 34, 34, 2, 241, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 
	2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 
	2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 
	2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 
	3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 
	77, 3, 2, 2, 2, 3, 81, 3, 2, 2, 2, 5, 83, 3, 2, 2, 2, 7, 86, 3, 2, 2, 2, 
	9, 89, 3, 2, 2, 2, 11, 92, 3, 2, 2, 2, 13, 95, 3, 2, 2, 2, 15, 97, 3, 2, 
	2, 2, 17, 99, 3, 2, 2, 2, 19, 102, 3, 2, 2, 2, 21, 105, 3, 2, 2, 2, 23, 
	107, 3, 2, 2, 2, 25, 109, 3, 2, 2, 2, 27, 111, 3, 2, 2, 2, 29, 113, 3, 
	2, 2, 2, 31, 115, 3, 2, 2, 2, 33, 117, 3, 2, 2, 2, 35, 119, 3, 2, 2, 2, 
	37, 121, 3, 2, 2, 2, 39, 126, 3, 2, 2, 2, 41, 132, 3, 2, 2, 2, 43, 136, 
	3, 2, 2, 2, 45, 139, 3, 2, 2, 2, 47, 144, 3, 2, 2, 2, 49, 150, 3, 2, 2, 
	2, 51, 154, 3, 2, 2, 2, 53, 160, 3, 2, 2, 2, 55, 165, 3, 2, 2, 2, 57, 173, 
	3, 2, 2, 2, 59, 195, 3, 2, 2, 2, 61, 200, 3, 2, 2, 2, 63, 210, 3, 2, 2, 
	2, 65, 214, 3, 2, 2, 2, 67, 216, 3, 2, 2, 2, 69, 218, 3, 2, 2, 2, 71, 220, 
	3, 2, 2, 2, 73, 222, 3, 2, 2, 2, 75, 224, 3, 2, 2, 2, 77, 226, 3, 2, 2, 
	2, 79, 82, 5, 57, 29, 2, 80, 82, 5, 59, 30, 2, 81, 79, 3, 2, 2, 2, 81, 
	80, 3, 2, 2, 2, 82, 4, 3, 2, 2, 2, 83, 84, 7, 126, 2, 2, 84, 85, 7, 126, 
	2, 2, 85, 6, 3, 2, 2, 2, 86, 87, 7, 40, 2, 2, 87, 88, 7, 40, 2, 2, 88, 
	8, 3, 2, 2, 2, 89, 90, 7, 63, 2, 2, 90, 91, 7, 63, 2, 2, 91, 10, 3, 2, 
	2, 2, 92, 93, 7, 35, 2, 2, 93, 94, 7, 63, 2, 2, 94, 12, 3, 2, 2, 2, 95, 
	96, 7, 64, 2, 2, 96, 14, 3, 2, 2, 2, 97, 98, 7, 62, 2, 2, 98, 16, 3, 2, 
	2, 2, 99, 100, 7, 64, 2, 2, 100, 101, 7, 63, 2, 2, 101, 18, 3, 2, 2, 2, 
	102, 103, 7, 62, 2, 2, 103, 104, 7, 63, 2, 2, 104, 20, 3, 2, 2, 2, 105, 
	106, 7, 45, 2, 2, 106, 22, 3, 2, 2, 2, 107, 108, 7, 47, 2, 2, 108, 24, 
	3, 2, 2, 2, 109, 110, 7, 44, 2, 2, 110, 26, 3, 2, 2, 2, 111, 112, 7, 49, 
	2, 2, 112, 28, 3, 2, 2, 2, 113, 114, 7, 39, 2, 2, 114, 30, 3, 2, 2, 2, 
	115, 116, 7, 35, 2, 2, 116, 32, 3, 2, 2, 2, 117, 118, 7, 60, 2, 2, 118, 
	34, 3, 2, 2, 2, 119, 120, 7, 61, 2, 2, 120, 36, 3, 2, 2, 2, 121, 122, 7, 
	118, 2, 2, 122, 123, 7, 116, 2, 2, 123, 124, 7, 119, 2, 2, 124, 125, 7, 
	103, 2, 2, 125, 38, 3, 2, 2, 2, 126, 127, 7, 104, 2, 2, 127, 128, 7, 99, 
	2, 2, 128, 129, 7, 110, 2, 2, 129, 130, 7, 117, 2, 2, 130, 131, 7, 103, 
	2, 2, 131, 40, 3, 2, 2, 2, 132, 133, 7, 104, 2, 2, 133, 134, 7, 113, 2, 
	2, 134, 135, 7, 116, 2, 2, 135, 42, 3, 2, 2, 2, 136, 137, 7, 107, 2, 2, 
	137, 138, 7, 104, 2, 2, 138, 44, 3, 2, 2, 2, 139, 140, 7, 103, 2, 2, 140, 
	141, 7, 110, 2, 2, 141, 142, 7, 117, 2, 2, 142, 143, 7, 103, 2, 2, 143, 
	46, 3, 2, 2, 2, 144, 145, 7, 121, 2, 2, 145, 146, 7, 106, 2, 2, 146, 147, 
	7, 107, 2, 2, 147, 148, 7, 110, 2, 2, 148, 149, 7, 103, 2, 2, 149, 48, 
	3, 2, 2, 2, 150, 151, 7, 107, 2, 2, 151, 152, 7, 112, 2, 2, 152, 153, 7, 
	118, 2, 2, 153, 50, 3, 2, 2, 2, 154, 155, 7, 104, 2, 2, 155, 156, 7, 110, 
	2, 2, 156, 157, 7, 113, 2, 2, 157, 158, 7, 99, 2, 2, 158, 159, 7, 118, 
	2, 2, 159, 52, 3, 2, 2, 2, 160, 161, 7, 100, 2, 2, 161, 162, 7, 113, 2, 
	2, 162, 163, 7, 113, 2, 2, 163, 164, 7, 110, 2, 2, 164, 54, 3, 2, 2, 2, 
	165, 169, 9, 2, 2, 2, 166, 168, 9, 3, 2, 2, 167, 166, 3, 2, 2, 2, 168, 
	171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 56, 3, 
	2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 174, 9, 4, 2, 2, 173, 172, 3, 2, 2, 
	2, 174, 175, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 
	58, 3, 2, 2, 2, 177, 179, 9, 4, 2, 2, 178, 177, 3, 2, 2, 2, 179, 180, 3, 
	2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 
	2, 182, 186, 7, 48, 2, 2, 183, 185, 9, 4, 2, 2, 184, 183, 3, 2, 2, 2, 185, 
	188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 196, 
	3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 191, 7, 48, 2, 2, 190, 192, 9, 4, 
	2, 2, 191, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 
	193, 194, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 178, 3, 2, 2, 2, 195, 
	189, 3, 2, 2, 2, 196, 60, 3, 2, 2, 2, 197, 201, 7, 37, 2, 2, 198, 199, 
	7, 49, 2, 2, 199, 201, 7, 49, 2, 2, 200, 197, 3, 2, 2, 2, 200, 198, 3, 
	2, 2, 2, 201, 205, 3, 2, 2, 2, 202, 204, 10, 5, 2, 2, 203, 202, 3, 2, 2, 
	2, 204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 
	208, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 209, 8, 31, 2, 2, 209, 62, 
	3, 2, 2, 2, 210, 211, 9, 6, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 8, 32, 
	2, 2, 213, 64, 3, 2, 2, 2, 214, 215, 7, 63, 2, 2, 215, 66, 3, 2, 2, 2, 
	216, 217, 7, 46, 2, 2, 217, 68, 3, 2, 2, 2, 218, 219, 7, 42, 2, 2, 219, 
	70, 3, 2, 2, 2, 220, 221, 7, 43, 2, 2, 221, 72, 3, 2, 2, 2, 222, 223, 7, 
	125, 2, 2, 223, 74, 3, 2, 2, 2, 224, 225, 7, 127, 2, 2, 225, 76, 3, 2, 
	2, 2, 226, 227, 7, 116, 2, 2, 227, 228, 7, 103, 2, 2, 228, 229, 7, 118, 
	2, 2, 229, 230, 7, 119, 2, 2, 230, 231, 7, 116, 2, 2, 231, 232, 7, 112, 
	2, 2, 232, 78, 3, 2, 2, 2, 12, 2, 81, 169, 175, 180, 186, 193, 195, 200, 
	205, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'", 
	"'-'", "'*'", "'/'", "'%'", "'!'", "':'", "';'", "'true'", "'false'", "'for'", 
	"'if'", "'else'", "'while'", "'int'", "'float'", "'bool'", "", "", "", 
	"", "", "'='", "','", "'('", "')'", "'{'", "'}'", "'return'",
}

var lexerSymbolicNames = []string{
	"", "NUMBER", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", 
	"MINUS", "MULT", "DIV", "MOD", "NOT", "COL", "SCOL", "TRUE", "FALSE", "FOR", 
	"IF", "ELSE", "WHILE", "INTTYPE", "FLOATTYPE", "BOOLTYPE", "ID", "INT", 
	"FLOAT", "COMMENT", "SPACE", "ASSIGN", "COM", "OPAR", "CPAR", "OBRACE", 
	"CBRACE", "RETURN",
}

var lexerRuleNames = []string{
	"NUMBER", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", 
	"MINUS", "MULT", "DIV", "MOD", "NOT", "COL", "SCOL", "TRUE", "FALSE", "FOR", 
	"IF", "ELSE", "WHILE", "INTTYPE", "FLOATTYPE", "BOOLTYPE", "ID", "INT", 
	"FLOAT", "COMMENT", "SPACE", "ASSIGN", "COM", "OPAR", "CPAR", "OBRACE", 
	"CBRACE", "RETURN",
}

type CLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCLexer(input antlr.CharStream) *CLexer {

	l := new(CLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "C.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CLexer tokens.
const (
	CLexerNUMBER = 1
	CLexerOR = 2
	CLexerAND = 3
	CLexerEQ = 4
	CLexerNEQ = 5
	CLexerGT = 6
	CLexerLT = 7
	CLexerGTEQ = 8
	CLexerLTEQ = 9
	CLexerPLUS = 10
	CLexerMINUS = 11
	CLexerMULT = 12
	CLexerDIV = 13
	CLexerMOD = 14
	CLexerNOT = 15
	CLexerCOL = 16
	CLexerSCOL = 17
	CLexerTRUE = 18
	CLexerFALSE = 19
	CLexerFOR = 20
	CLexerIF = 21
	CLexerELSE = 22
	CLexerWHILE = 23
	CLexerINTTYPE = 24
	CLexerFLOATTYPE = 25
	CLexerBOOLTYPE = 26
	CLexerID = 27
	CLexerINT = 28
	CLexerFLOAT = 29
	CLexerCOMMENT = 30
	CLexerSPACE = 31
	CLexerASSIGN = 32
	CLexerCOM = 33
	CLexerOPAR = 34
	CLexerCPAR = 35
	CLexerOBRACE = 36
	CLexerCBRACE = 37
	CLexerRETURN = 38
)

