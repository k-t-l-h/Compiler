grammar C;

function: typeId  ID OPAR  argList CPAR compoundStatement EOF;
declaration: typeId  ID ;
argList: arg | arg COM  argList|;

arg: typeId ID;

typeId:  INTTYPE | FLOATTYPE | BOOLTYPE;

iDList: ID | ID COM iDList;

compoundStatement: OBRACE  statement+ CBRACE  | OBRACE  statement+ returnStatement CBRACE ;
returnStatement: RETURN expression SCOL;

statementList: statement+ |;

statement: forStatement |
          whileStatement|
          ifStatement|
          compoundStatement|
          expression SCOL |
          declaration SCOL;

forStatement: FOR OPAR expression SCOL optionExpression SCOL expression CPAR compoundStatement;
optionExpression: expression |;
whileStatement: WHILE OPAR expression CPAR compoundStatement;
ifStatement: IF ( expression ) compoundStatement elseStatement;
elseStatement: ELSE compoundStatement |;

expression: ID ASSIGN expression | ID ASSIGN rValue | ID ASSIGN term;
rValue: magnitude compare magnitude | magnitude | magnitude OR magnitude | magnitude AND magnitude;
compare: ('=='| '<' | '<=' | '>' | '>=' | '!=');
magnitude: term  PLUS magnitude |  term  MINUS magnitude | term MOD factor| term MULT factor
| term DIV factor | factor;
factor: OPAR expression CPAR | MINUS factor | PLUS factor |  NOT factor;
term: ID | NUMBER;


NUMBER: INT | FLOAT;
OR : '||';
AND : '&&';
EQ : '==';
NEQ : '!=';
GT : '>';
LT : '<';
GTEQ : '>=';
LTEQ : '<=';
PLUS : '+';
MINUS : '-';
MULT : '*';
DIV : '/';
MOD : '%';
NOT : '!';

COL: ':';
SCOL : ';';

TRUE : 'true';
FALSE : 'false';
FOR: 'for';
IF : 'if';
ELSE : 'else';
WHILE : 'while';

INTTYPE: 'int';
FLOATTYPE: 'float';
BOOLTYPE : 'bool';

ID
 : [a-zA-Z_] [a-zA-Z_0-9]*
 ;

INT
 : [0-9]+
 ;

FLOAT
 : [0-9]+ '.' [0-9]*
 | '.' [0-9]+
 ;

COMMENT: ('#' | '//') ~[\r\n]* -> skip
 ;

SPACE
 : [ \t\r\n] -> skip
 ;

ASSIGN : '=';
COM : ',';
OPAR : '(';
CPAR : ')';
OBRACE : '{';
CBRACE : '}';

RETURN : 'return';