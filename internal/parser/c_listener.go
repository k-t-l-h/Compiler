// Generated from C.g4 by ANTLR 4.7.

package parser // C

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CListener is a complete listener for a parse tree produced by CParser.
type CListener interface {
	antlr.ParseTreeListener

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterTypeId is called when entering the typeId production.
	EnterTypeId(c *TypeIdContext)

	// EnterIDList is called when entering the iDList production.
	EnterIDList(c *IDListContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterOptionExpression is called when entering the optionExpression production.
	EnterOptionExpression(c *OptionExpressionContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRValue is called when entering the rValue production.
	EnterRValue(c *RValueContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterMagnitude is called when entering the magnitude production.
	EnterMagnitude(c *MagnitudeContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitTypeId is called when exiting the typeId production.
	ExitTypeId(c *TypeIdContext)

	// ExitIDList is called when exiting the iDList production.
	ExitIDList(c *IDListContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitOptionExpression is called when exiting the optionExpression production.
	ExitOptionExpression(c *OptionExpressionContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRValue is called when exiting the rValue production.
	ExitRValue(c *RValueContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitMagnitude is called when exiting the magnitude production.
	ExitMagnitude(c *MagnitudeContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
