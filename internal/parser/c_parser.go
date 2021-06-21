// Generated from C.g4 by ANTLR 4.7.

package parser // C

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 219, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2, 
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 
	3, 4, 3, 4, 3, 4, 5, 4, 62, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 
	3, 7, 3, 7, 3, 7, 5, 7, 73, 10, 7, 3, 8, 3, 8, 6, 8, 77, 10, 8, 13, 8, 
	14, 8, 78, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 85, 10, 8, 13, 8, 14, 8, 86, 3, 
	8, 3, 8, 3, 8, 5, 8, 92, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 99, 
	10, 10, 13, 10, 14, 10, 100, 3, 10, 5, 10, 104, 10, 10, 3, 11, 3, 11, 3, 
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 116, 10, 11, 
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 
	13, 3, 13, 5, 13, 130, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 
	16, 148, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 
	3, 17, 5, 17, 159, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 
	5, 18, 178, 10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 203, 10, 20, 3, 21, 3, 21, 3, 
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 215, 10, 21, 
	3, 22, 3, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 5, 3, 2, 26, 28, 3, 2, 6, 11, 
	4, 2, 3, 3, 29, 29, 2, 226, 2, 44, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 6, 61, 
	3, 2, 2, 2, 8, 63, 3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 
	14, 91, 3, 2, 2, 2, 16, 93, 3, 2, 2, 2, 18, 103, 3, 2, 2, 2, 20, 115, 3, 
	2, 2, 2, 22, 117, 3, 2, 2, 2, 24, 129, 3, 2, 2, 2, 26, 131, 3, 2, 2, 2, 
	28, 137, 3, 2, 2, 2, 30, 147, 3, 2, 2, 2, 32, 158, 3, 2, 2, 2, 34, 177, 
	3, 2, 2, 2, 36, 179, 3, 2, 2, 2, 38, 202, 3, 2, 2, 2, 40, 214, 3, 2, 2, 
	2, 42, 216, 3, 2, 2, 2, 44, 45, 5, 10, 6, 2, 45, 46, 7, 29, 2, 2, 46, 47, 
	7, 36, 2, 2, 47, 48, 5, 6, 4, 2, 48, 49, 7, 37, 2, 2, 49, 50, 5, 14, 8, 
	2, 50, 51, 7, 2, 2, 3, 51, 3, 3, 2, 2, 2, 52, 53, 5, 10, 6, 2, 53, 54, 
	7, 29, 2, 2, 54, 5, 3, 2, 2, 2, 55, 62, 5, 8, 5, 2, 56, 57, 5, 8, 5, 2, 
	57, 58, 7, 35, 2, 2, 58, 59, 5, 6, 4, 2, 59, 62, 3, 2, 2, 2, 60, 62, 3, 
	2, 2, 2, 61, 55, 3, 2, 2, 2, 61, 56, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 
	7, 3, 2, 2, 2, 63, 64, 5, 10, 6, 2, 64, 65, 7, 29, 2, 2, 65, 9, 3, 2, 2, 
	2, 66, 67, 9, 2, 2, 2, 67, 11, 3, 2, 2, 2, 68, 73, 7, 29, 2, 2, 69, 70, 
	7, 29, 2, 2, 70, 71, 7, 35, 2, 2, 71, 73, 5, 12, 7, 2, 72, 68, 3, 2, 2, 
	2, 72, 69, 3, 2, 2, 2, 73, 13, 3, 2, 2, 2, 74, 76, 7, 38, 2, 2, 75, 77, 
	5, 20, 11, 2, 76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 
	2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 7, 39, 2, 2, 81, 92, 
	3, 2, 2, 2, 82, 84, 7, 38, 2, 2, 83, 85, 5, 20, 11, 2, 84, 83, 3, 2, 2, 
	2, 85, 86, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 
	3, 2, 2, 2, 88, 89, 5, 16, 9, 2, 89, 90, 7, 39, 2, 2, 90, 92, 3, 2, 2, 
	2, 91, 74, 3, 2, 2, 2, 91, 82, 3, 2, 2, 2, 92, 15, 3, 2, 2, 2, 93, 94, 
	7, 40, 2, 2, 94, 95, 5, 32, 17, 2, 95, 96, 7, 19, 2, 2, 96, 17, 3, 2, 2, 
	2, 97, 99, 5, 20, 11, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 
	98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 104, 3, 
	2, 2, 2, 103, 98, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 19, 3, 2, 2, 2, 
	105, 116, 5, 22, 12, 2, 106, 116, 5, 26, 14, 2, 107, 116, 5, 28, 15, 2, 
	108, 116, 5, 14, 8, 2, 109, 110, 5, 32, 17, 2, 110, 111, 7, 19, 2, 2, 111, 
	116, 3, 2, 2, 2, 112, 113, 5, 4, 3, 2, 113, 114, 7, 19, 2, 2, 114, 116, 
	3, 2, 2, 2, 115, 105, 3, 2, 2, 2, 115, 106, 3, 2, 2, 2, 115, 107, 3, 2, 
	2, 2, 115, 108, 3, 2, 2, 2, 115, 109, 3, 2, 2, 2, 115, 112, 3, 2, 2, 2, 
	116, 21, 3, 2, 2, 2, 117, 118, 7, 22, 2, 2, 118, 119, 7, 36, 2, 2, 119, 
	120, 5, 32, 17, 2, 120, 121, 7, 19, 2, 2, 121, 122, 5, 24, 13, 2, 122, 
	123, 7, 19, 2, 2, 123, 124, 5, 32, 17, 2, 124, 125, 7, 37, 2, 2, 125, 126, 
	5, 14, 8, 2, 126, 23, 3, 2, 2, 2, 127, 130, 5, 32, 17, 2, 128, 130, 3, 
	2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 25, 3, 2, 2, 
	2, 131, 132, 7, 25, 2, 2, 132, 133, 7, 36, 2, 2, 133, 134, 5, 34, 18, 2, 
	134, 135, 7, 37, 2, 2, 135, 136, 5, 14, 8, 2, 136, 27, 3, 2, 2, 2, 137, 
	138, 7, 23, 2, 2, 138, 139, 7, 36, 2, 2, 139, 140, 5, 34, 18, 2, 140, 141, 
	7, 37, 2, 2, 141, 142, 5, 14, 8, 2, 142, 143, 5, 30, 16, 2, 143, 29, 3, 
	2, 2, 2, 144, 145, 7, 24, 2, 2, 145, 148, 5, 14, 8, 2, 146, 148, 3, 2, 
	2, 2, 147, 144, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 31, 3, 2, 2, 2, 
	149, 150, 7, 29, 2, 2, 150, 151, 7, 34, 2, 2, 151, 159, 5, 32, 17, 2, 152, 
	153, 7, 29, 2, 2, 153, 154, 7, 34, 2, 2, 154, 159, 5, 34, 18, 2, 155, 156, 
	7, 29, 2, 2, 156, 157, 7, 34, 2, 2, 157, 159, 5, 42, 22, 2, 158, 149, 3, 
	2, 2, 2, 158, 152, 3, 2, 2, 2, 158, 155, 3, 2, 2, 2, 159, 33, 3, 2, 2, 
	2, 160, 161, 5, 38, 20, 2, 161, 162, 5, 36, 19, 2, 162, 163, 5, 38, 20, 
	2, 163, 178, 3, 2, 2, 2, 164, 178, 5, 38, 20, 2, 165, 166, 5, 38, 20, 2, 
	166, 167, 7, 4, 2, 2, 167, 168, 5, 38, 20, 2, 168, 178, 3, 2, 2, 2, 169, 
	170, 5, 38, 20, 2, 170, 171, 7, 5, 2, 2, 171, 172, 5, 38, 20, 2, 172, 178, 
	3, 2, 2, 2, 173, 174, 5, 42, 22, 2, 174, 175, 5, 36, 19, 2, 175, 176, 5, 
	42, 22, 2, 176, 178, 3, 2, 2, 2, 177, 160, 3, 2, 2, 2, 177, 164, 3, 2, 
	2, 2, 177, 165, 3, 2, 2, 2, 177, 169, 3, 2, 2, 2, 177, 173, 3, 2, 2, 2, 
	178, 35, 3, 2, 2, 2, 179, 180, 9, 3, 2, 2, 180, 37, 3, 2, 2, 2, 181, 182, 
	5, 42, 22, 2, 182, 183, 7, 12, 2, 2, 183, 184, 5, 38, 20, 2, 184, 203, 
	3, 2, 2, 2, 185, 186, 5, 42, 22, 2, 186, 187, 7, 13, 2, 2, 187, 188, 5, 
	38, 20, 2, 188, 203, 3, 2, 2, 2, 189, 190, 5, 42, 22, 2, 190, 191, 7, 16, 
	2, 2, 191, 192, 5, 40, 21, 2, 192, 203, 3, 2, 2, 2, 193, 194, 5, 42, 22, 
	2, 194, 195, 7, 14, 2, 2, 195, 196, 5, 40, 21, 2, 196, 203, 3, 2, 2, 2, 
	197, 198, 5, 42, 22, 2, 198, 199, 7, 15, 2, 2, 199, 200, 5, 40, 21, 2, 
	200, 203, 3, 2, 2, 2, 201, 203, 5, 40, 21, 2, 202, 181, 3, 2, 2, 2, 202, 
	185, 3, 2, 2, 2, 202, 189, 3, 2, 2, 2, 202, 193, 3, 2, 2, 2, 202, 197, 
	3, 2, 2, 2, 202, 201, 3, 2, 2, 2, 203, 39, 3, 2, 2, 2, 204, 205, 7, 36, 
	2, 2, 205, 206, 5, 32, 17, 2, 206, 207, 7, 37, 2, 2, 207, 215, 3, 2, 2, 
	2, 208, 209, 7, 13, 2, 2, 209, 215, 5, 40, 21, 2, 210, 211, 7, 12, 2, 2, 
	211, 215, 5, 40, 21, 2, 212, 213, 7, 17, 2, 2, 213, 215, 5, 40, 21, 2, 
	214, 204, 3, 2, 2, 2, 214, 208, 3, 2, 2, 2, 214, 210, 3, 2, 2, 2, 214, 
	212, 3, 2, 2, 2, 215, 41, 3, 2, 2, 2, 216, 217, 9, 4, 2, 2, 217, 43, 3, 
	2, 2, 2, 16, 61, 72, 78, 86, 91, 100, 103, 115, 129, 147, 158, 177, 202, 
	214,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'", 
	"'-'", "'*'", "'/'", "'%'", "'!'", "':'", "';'", "'true'", "'false'", "'for'", 
	"'if'", "'else'", "'while'", "'int'", "'float'", "'bool'", "", "", "", 
	"", "", "'='", "','", "'('", "')'", "'{'", "'}'", "'return'",
}
var symbolicNames = []string{
	"", "NUMBER", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", 
	"MINUS", "MULT", "DIV", "MOD", "NOT", "COL", "SCOL", "TRUE", "FALSE", "FOR", 
	"IF", "ELSE", "WHILE", "INTTYPE", "FLOATTYPE", "BOOLTYPE", "ID", "INT", 
	"FLOAT", "COMMENT", "SPACE", "ASSIGN", "COM", "OPAR", "CPAR", "OBRACE", 
	"CBRACE", "RETURN",
}

var ruleNames = []string{
	"function", "declaration", "argList", "arg", "typeId", "iDList", "compoundStatement", 
	"returnStatement", "statementList", "statement", "forStatement", "optionExpression", 
	"whileStatement", "ifStatement", "elseStatement", "expression", "rValue", 
	"compare", "magnitude", "factor", "term",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CParser struct {
	*antlr.BaseParser
}

func NewCParser(input antlr.TokenStream) *CParser {
	this := new(CParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "C.g4"

	return this
}

// CParser tokens.
const (
	CParserEOF = antlr.TokenEOF
	CParserNUMBER = 1
	CParserOR = 2
	CParserAND = 3
	CParserEQ = 4
	CParserNEQ = 5
	CParserGT = 6
	CParserLT = 7
	CParserGTEQ = 8
	CParserLTEQ = 9
	CParserPLUS = 10
	CParserMINUS = 11
	CParserMULT = 12
	CParserDIV = 13
	CParserMOD = 14
	CParserNOT = 15
	CParserCOL = 16
	CParserSCOL = 17
	CParserTRUE = 18
	CParserFALSE = 19
	CParserFOR = 20
	CParserIF = 21
	CParserELSE = 22
	CParserWHILE = 23
	CParserINTTYPE = 24
	CParserFLOATTYPE = 25
	CParserBOOLTYPE = 26
	CParserID = 27
	CParserINT = 28
	CParserFLOAT = 29
	CParserCOMMENT = 30
	CParserSPACE = 31
	CParserASSIGN = 32
	CParserCOM = 33
	CParserOPAR = 34
	CParserCPAR = 35
	CParserOBRACE = 36
	CParserCBRACE = 37
	CParserRETURN = 38
)

// CParser rules.
const (
	CParserRULE_function = 0
	CParserRULE_declaration = 1
	CParserRULE_argList = 2
	CParserRULE_arg = 3
	CParserRULE_typeId = 4
	CParserRULE_iDList = 5
	CParserRULE_compoundStatement = 6
	CParserRULE_returnStatement = 7
	CParserRULE_statementList = 8
	CParserRULE_statement = 9
	CParserRULE_forStatement = 10
	CParserRULE_optionExpression = 11
	CParserRULE_whileStatement = 12
	CParserRULE_ifStatement = 13
	CParserRULE_elseStatement = 14
	CParserRULE_expression = 15
	CParserRULE_rValue = 16
	CParserRULE_compare = 17
	CParserRULE_magnitude = 18
	CParserRULE_factor = 19
	CParserRULE_term = 20
)

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) TypeId() ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *FunctionContext) OPAR() antlr.TerminalNode {
	return s.GetToken(CParserOPAR, 0)
}

func (s *FunctionContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionContext) CPAR() antlr.TerminalNode {
	return s.GetToken(CParserCPAR, 0)
}

func (s *FunctionContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *FunctionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CParserEOF, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitFunction(s)
	}
}




func (p *CParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.TypeId()
	}
	{
		p.SetState(43)
		p.Match(CParserID)
	}
	{
		p.SetState(44)
		p.Match(CParserOPAR)
	}
	{
		p.SetState(45)
		p.ArgList()
	}
	{
		p.SetState(46)
		p.Match(CParserCPAR)
	}
	{
		p.SetState(47)
		p.CompoundStatement()
	}
	{
		p.SetState(48)
		p.Match(CParserEOF)
	}



	return localctx
}


// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeId() ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitDeclaration(s)
	}
}




func (p *CParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.TypeId()
	}
	{
		p.SetState(51)
		p.Match(CParserID)
	}



	return localctx
}


// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) Arg() IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ArgListContext) COM() antlr.TerminalNode {
	return s.GetToken(CParserCOM, 0)
}

func (s *ArgListContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitArgList(s)
	}
}




func (p *CParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CParserRULE_argList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Arg()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Arg()
		}
		{
			p.SetState(55)
			p.Match(CParserCOM)
		}
		{
			p.SetState(56)
			p.ArgList()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)

	}


	return localctx
}


// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) TypeId() ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *ArgContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitArg(s)
	}
}




func (p *CParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CParserRULE_arg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.TypeId()
	}
	{
		p.SetState(62)
		p.Match(CParserID)
	}



	return localctx
}


// ITypeIdContext is an interface to support dynamic dispatch.
type ITypeIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeIdContext differentiates from other interfaces.
	IsTypeIdContext()
}

type TypeIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeIdContext() *TypeIdContext {
	var p = new(TypeIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_typeId
	return p
}

func (*TypeIdContext) IsTypeIdContext() {}

func NewTypeIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeIdContext {
	var p = new(TypeIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_typeId

	return p
}

func (s *TypeIdContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeIdContext) INTTYPE() antlr.TerminalNode {
	return s.GetToken(CParserINTTYPE, 0)
}

func (s *TypeIdContext) FLOATTYPE() antlr.TerminalNode {
	return s.GetToken(CParserFLOATTYPE, 0)
}

func (s *TypeIdContext) BOOLTYPE() antlr.TerminalNode {
	return s.GetToken(CParserBOOLTYPE, 0)
}

func (s *TypeIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterTypeId(s)
	}
}

func (s *TypeIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitTypeId(s)
	}
}




func (p *CParser) TypeId() (localctx ITypeIdContext) {
	localctx = NewTypeIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CParserRULE_typeId)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << CParserINTTYPE) | (1 << CParserFLOATTYPE) | (1 << CParserBOOLTYPE))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IIDListContext is an interface to support dynamic dispatch.
type IIDListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIDListContext differentiates from other interfaces.
	IsIDListContext()
}

type IDListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIDListContext() *IDListContext {
	var p = new(IDListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_iDList
	return p
}

func (*IDListContext) IsIDListContext() {}

func NewIDListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IDListContext {
	var p = new(IDListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_iDList

	return p
}

func (s *IDListContext) GetParser() antlr.Parser { return s.parser }

func (s *IDListContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *IDListContext) COM() antlr.TerminalNode {
	return s.GetToken(CParserCOM, 0)
}

func (s *IDListContext) IDList() IIDListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIDListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIDListContext)
}

func (s *IDListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IDListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterIDList(s)
	}
}

func (s *IDListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitIDList(s)
	}
}




func (p *CParser) IDList() (localctx IIDListContext) {
	localctx = NewIDListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CParserRULE_iDList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(CParserID)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(CParserID)
		}
		{
			p.SetState(68)
			p.Match(CParserCOM)
		}
		{
			p.SetState(69)
			p.IDList()
		}

	}


	return localctx
}


// ICompoundStatementContext is an interface to support dynamic dispatch.
type ICompoundStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundStatementContext differentiates from other interfaces.
	IsCompoundStatementContext()
}

type CompoundStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundStatementContext() *CompoundStatementContext {
	var p = new(CompoundStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_compoundStatement
	return p
}

func (*CompoundStatementContext) IsCompoundStatementContext() {}

func NewCompoundStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundStatementContext {
	var p = new(CompoundStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_compoundStatement

	return p
}

func (s *CompoundStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundStatementContext) OBRACE() antlr.TerminalNode {
	return s.GetToken(CParserOBRACE, 0)
}

func (s *CompoundStatementContext) CBRACE() antlr.TerminalNode {
	return s.GetToken(CParserCBRACE, 0)
}

func (s *CompoundStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CompoundStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CompoundStatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *CompoundStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CompoundStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCompoundStatement(s)
	}
}




func (p *CParser) CompoundStatement() (localctx ICompoundStatementContext) {
	localctx = NewCompoundStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CParserRULE_compoundStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(CParserOBRACE)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((((_la - 20)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 20))) & ((1 << (CParserFOR - 20)) | (1 << (CParserIF - 20)) | (1 << (CParserWHILE - 20)) | (1 << (CParserINTTYPE - 20)) | (1 << (CParserFLOATTYPE - 20)) | (1 << (CParserBOOLTYPE - 20)) | (1 << (CParserID - 20)) | (1 << (CParserOBRACE - 20)))) != 0) {
			{
				p.SetState(73)
				p.Statement()
			}


			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(78)
			p.Match(CParserCBRACE)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(CParserOBRACE)
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((((_la - 20)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 20))) & ((1 << (CParserFOR - 20)) | (1 << (CParserIF - 20)) | (1 << (CParserWHILE - 20)) | (1 << (CParserINTTYPE - 20)) | (1 << (CParserFLOATTYPE - 20)) | (1 << (CParserBOOLTYPE - 20)) | (1 << (CParserID - 20)) | (1 << (CParserOBRACE - 20)))) != 0) {
			{
				p.SetState(81)
				p.Statement()
			}


			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(86)
			p.ReturnStatement()
		}
		{
			p.SetState(87)
			p.Match(CParserCBRACE)
		}

	}


	return localctx
}


// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) SCOL() antlr.TerminalNode {
	return s.GetToken(CParserSCOL, 0)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}




func (p *CParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(CParserRETURN)
	}
	{
		p.SetState(92)
		p.Expression()
	}
	{
		p.SetState(93)
		p.Match(CParserSCOL)
	}



	return localctx
}


// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitStatementList(s)
	}
}




func (p *CParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CParserRULE_statementList)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserFOR, CParserIF, CParserWHILE, CParserINTTYPE, CParserFLOATTYPE, CParserBOOLTYPE, CParserID, CParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((((_la - 20)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 20))) & ((1 << (CParserFOR - 20)) | (1 << (CParserIF - 20)) | (1 << (CParserWHILE - 20)) | (1 << (CParserINTTYPE - 20)) | (1 << (CParserFLOATTYPE - 20)) | (1 << (CParserBOOLTYPE - 20)) | (1 << (CParserID - 20)) | (1 << (CParserOBRACE - 20)))) != 0) {
			{
				p.SetState(95)
				p.Statement()
			}


			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case CParserEOF:
		p.EnterOuterAlt(localctx, 2)



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) SCOL() antlr.TerminalNode {
	return s.GetToken(CParserSCOL, 0)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitStatement(s)
	}
}




func (p *CParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserFOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.ForStatement()
		}


	case CParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.WhileStatement()
		}


	case CParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.IfStatement()
		}


	case CParserOBRACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.CompoundStatement()
		}


	case CParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Expression()
		}
		{
			p.SetState(108)
			p.Match(CParserSCOL)
		}


	case CParserINTTYPE, CParserFLOATTYPE, CParserBOOLTYPE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
			p.Declaration()
		}
		{
			p.SetState(111)
			p.Match(CParserSCOL)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(CParserFOR, 0)
}

func (s *ForStatementContext) OPAR() antlr.TerminalNode {
	return s.GetToken(CParserOPAR, 0)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) AllSCOL() []antlr.TerminalNode {
	return s.GetTokens(CParserSCOL)
}

func (s *ForStatementContext) SCOL(i int) antlr.TerminalNode {
	return s.GetToken(CParserSCOL, i)
}

func (s *ForStatementContext) OptionExpression() IOptionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionExpressionContext)
}

func (s *ForStatementContext) CPAR() antlr.TerminalNode {
	return s.GetToken(CParserCPAR, 0)
}

func (s *ForStatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitForStatement(s)
	}
}




func (p *CParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CParserRULE_forStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(CParserFOR)
	}
	{
		p.SetState(116)
		p.Match(CParserOPAR)
	}
	{
		p.SetState(117)
		p.Expression()
	}
	{
		p.SetState(118)
		p.Match(CParserSCOL)
	}
	{
		p.SetState(119)
		p.OptionExpression()
	}
	{
		p.SetState(120)
		p.Match(CParserSCOL)
	}
	{
		p.SetState(121)
		p.Expression()
	}
	{
		p.SetState(122)
		p.Match(CParserCPAR)
	}
	{
		p.SetState(123)
		p.CompoundStatement()
	}



	return localctx
}


// IOptionExpressionContext is an interface to support dynamic dispatch.
type IOptionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionExpressionContext differentiates from other interfaces.
	IsOptionExpressionContext()
}

type OptionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionExpressionContext() *OptionExpressionContext {
	var p = new(OptionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_optionExpression
	return p
}

func (*OptionExpressionContext) IsOptionExpressionContext() {}

func NewOptionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionExpressionContext {
	var p = new(OptionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_optionExpression

	return p
}

func (s *OptionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OptionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterOptionExpression(s)
	}
}

func (s *OptionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitOptionExpression(s)
	}
}




func (p *CParser) OptionExpression() (localctx IOptionExpressionContext) {
	localctx = NewOptionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CParserRULE_optionExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Expression()
		}


	case CParserSCOL:
		p.EnterOuterAlt(localctx, 2)



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(CParserWHILE, 0)
}

func (s *WhileStatementContext) OPAR() antlr.TerminalNode {
	return s.GetToken(CParserOPAR, 0)
}

func (s *WhileStatementContext) RValue() IRValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValueContext)
}

func (s *WhileStatementContext) CPAR() antlr.TerminalNode {
	return s.GetToken(CParserCPAR, 0)
}

func (s *WhileStatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}




func (p *CParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(CParserWHILE)
	}
	{
		p.SetState(130)
		p.Match(CParserOPAR)
	}
	{
		p.SetState(131)
		p.RValue()
	}
	{
		p.SetState(132)
		p.Match(CParserCPAR)
	}
	{
		p.SetState(133)
		p.CompoundStatement()
	}



	return localctx
}


// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(CParserIF, 0)
}

func (s *IfStatementContext) OPAR() antlr.TerminalNode {
	return s.GetToken(CParserOPAR, 0)
}

func (s *IfStatementContext) RValue() IRValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValueContext)
}

func (s *IfStatementContext) CPAR() antlr.TerminalNode {
	return s.GetToken(CParserCPAR, 0)
}

func (s *IfStatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitIfStatement(s)
	}
}




func (p *CParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(CParserIF)
	}
	{
		p.SetState(136)
		p.Match(CParserOPAR)
	}
	{
		p.SetState(137)
		p.RValue()
	}
	{
		p.SetState(138)
		p.Match(CParserCPAR)
	}
	{
		p.SetState(139)
		p.CompoundStatement()
	}
	{
		p.SetState(140)
		p.ElseStatement()
	}



	return localctx
}


// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CParserELSE, 0)
}

func (s *ElseStatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitElseStatement(s)
	}
}




func (p *CParser) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CParserRULE_elseStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserELSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(CParserELSE)
		}
		{
			p.SetState(143)
			p.CompoundStatement()
		}


	case CParserEOF, CParserFOR, CParserIF, CParserWHILE, CParserINTTYPE, CParserFLOATTYPE, CParserBOOLTYPE, CParserID, CParserOBRACE, CParserCBRACE, CParserRETURN:
		p.EnterOuterAlt(localctx, 2)



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CParserASSIGN, 0)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RValue() IRValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValueContext)
}

func (s *ExpressionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitExpression(s)
	}
}




func (p *CParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Match(CParserID)
		}
		{
			p.SetState(148)
			p.Match(CParserASSIGN)
		}
		{
			p.SetState(149)
			p.Expression()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(CParserID)
		}
		{
			p.SetState(151)
			p.Match(CParserASSIGN)
		}
		{
			p.SetState(152)
			p.RValue()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.Match(CParserID)
		}
		{
			p.SetState(154)
			p.Match(CParserASSIGN)
		}
		{
			p.SetState(155)
			p.Term()
		}

	}


	return localctx
}


// IRValueContext is an interface to support dynamic dispatch.
type IRValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRValueContext differentiates from other interfaces.
	IsRValueContext()
}

type RValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRValueContext() *RValueContext {
	var p = new(RValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_rValue
	return p
}

func (*RValueContext) IsRValueContext() {}

func NewRValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RValueContext {
	var p = new(RValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_rValue

	return p
}

func (s *RValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RValueContext) AllMagnitude() []IMagnitudeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMagnitudeContext)(nil)).Elem())
	var tst = make([]IMagnitudeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMagnitudeContext)
		}
	}

	return tst
}

func (s *RValueContext) Magnitude(i int) IMagnitudeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMagnitudeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMagnitudeContext)
}

func (s *RValueContext) Compare() ICompareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *RValueContext) OR() antlr.TerminalNode {
	return s.GetToken(CParserOR, 0)
}

func (s *RValueContext) AND() antlr.TerminalNode {
	return s.GetToken(CParserAND, 0)
}

func (s *RValueContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *RValueContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *RValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterRValue(s)
	}
}

func (s *RValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitRValue(s)
	}
}




func (p *CParser) RValue() (localctx IRValueContext) {
	localctx = NewRValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CParserRULE_rValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Magnitude()
		}
		{
			p.SetState(159)
			p.Compare()
		}
		{
			p.SetState(160)
			p.Magnitude()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Magnitude()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Magnitude()
		}
		{
			p.SetState(164)
			p.Match(CParserOR)
		}
		{
			p.SetState(165)
			p.Magnitude()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(167)
			p.Magnitude()
		}
		{
			p.SetState(168)
			p.Match(CParserAND)
		}
		{
			p.SetState(169)
			p.Magnitude()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(171)
			p.Term()
		}
		{
			p.SetState(172)
			p.Compare()
		}
		{
			p.SetState(173)
			p.Term()
		}

	}


	return localctx
}


// ICompareContext is an interface to support dynamic dispatch.
type ICompareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareContext differentiates from other interfaces.
	IsCompareContext()
}

type CompareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareContext() *CompareContext {
	var p = new(CompareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_compare
	return p
}

func (*CompareContext) IsCompareContext() {}

func NewCompareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareContext {
	var p = new(CompareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_compare

	return p
}

func (s *CompareContext) GetParser() antlr.Parser { return s.parser }
func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitCompare(s)
	}
}




func (p *CParser) Compare() (localctx ICompareContext) {
	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CParserRULE_compare)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << CParserEQ) | (1 << CParserNEQ) | (1 << CParserGT) | (1 << CParserLT) | (1 << CParserGTEQ) | (1 << CParserLTEQ))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IMagnitudeContext is an interface to support dynamic dispatch.
type IMagnitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMagnitudeContext differentiates from other interfaces.
	IsMagnitudeContext()
}

type MagnitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMagnitudeContext() *MagnitudeContext {
	var p = new(MagnitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_magnitude
	return p
}

func (*MagnitudeContext) IsMagnitudeContext() {}

func NewMagnitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MagnitudeContext {
	var p = new(MagnitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_magnitude

	return p
}

func (s *MagnitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *MagnitudeContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *MagnitudeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CParserPLUS, 0)
}

func (s *MagnitudeContext) Magnitude() IMagnitudeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMagnitudeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMagnitudeContext)
}

func (s *MagnitudeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CParserMINUS, 0)
}

func (s *MagnitudeContext) MOD() antlr.TerminalNode {
	return s.GetToken(CParserMOD, 0)
}

func (s *MagnitudeContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *MagnitudeContext) MULT() antlr.TerminalNode {
	return s.GetToken(CParserMULT, 0)
}

func (s *MagnitudeContext) DIV() antlr.TerminalNode {
	return s.GetToken(CParserDIV, 0)
}

func (s *MagnitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MagnitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MagnitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterMagnitude(s)
	}
}

func (s *MagnitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitMagnitude(s)
	}
}




func (p *CParser) Magnitude() (localctx IMagnitudeContext) {
	localctx = NewMagnitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CParserRULE_magnitude)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Term()
		}
		{
			p.SetState(180)
			p.Match(CParserPLUS)
		}
		{
			p.SetState(181)
			p.Magnitude()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Term()
		}
		{
			p.SetState(184)
			p.Match(CParserMINUS)
		}
		{
			p.SetState(185)
			p.Magnitude()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Term()
		}
		{
			p.SetState(188)
			p.Match(CParserMOD)
		}
		{
			p.SetState(189)
			p.Factor()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(191)
			p.Term()
		}
		{
			p.SetState(192)
			p.Match(CParserMULT)
		}
		{
			p.SetState(193)
			p.Factor()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.Term()
		}
		{
			p.SetState(196)
			p.Match(CParserDIV)
		}
		{
			p.SetState(197)
			p.Factor()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(199)
			p.Factor()
		}

	}


	return localctx
}


// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) OPAR() antlr.TerminalNode {
	return s.GetToken(CParserOPAR, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) CPAR() antlr.TerminalNode {
	return s.GetToken(CParserCPAR, 0)
}

func (s *FactorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CParserMINUS, 0)
}

func (s *FactorContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CParserPLUS, 0)
}

func (s *FactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(CParserNOT, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitFactor(s)
	}
}




func (p *CParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserOPAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(CParserOPAR)
		}
		{
			p.SetState(203)
			p.Expression()
		}
		{
			p.SetState(204)
			p.Match(CParserCPAR)
		}


	case CParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(CParserMINUS)
		}
		{
			p.SetState(207)
			p.Factor()
		}


	case CParserPLUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.Match(CParserPLUS)
		}
		{
			p.SetState(209)
			p.Factor()
		}


	case CParserNOT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)
			p.Match(CParserNOT)
		}
		{
			p.SetState(211)
			p.Factor()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) ID() antlr.TerminalNode {
	return s.GetToken(CParserID, 0)
}

func (s *TermContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CParserNUMBER, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CListener); ok {
		listenerT.ExitTerm(s)
	}
}




func (p *CParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CParserRULE_term)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(214)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CParserNUMBER || _la == CParserID) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


