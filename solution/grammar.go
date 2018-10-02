// Code generated by goyacc -o grammar.go grammar.y. DO NOT EDIT.

//line grammar.y:2

package solution

import __yyfmt__ "fmt"

//line grammar.y:4

//line grammar.y:6
type yySymType struct {
	yys  int
	tok  tokenType
	str  string
	line int
}

const COMMA = 57346
const DOT = 57347
const EQ = 57348
const BRACE_OPEN = 57349
const BRACE_CLOSE = 57350
const PAREN_OPEN = 57351
const PAREN_CLOSE = 57352
const COMMENT = 57353
const NUMBER = 57354
const CRLF = 57355
const IDENTIFIER = 57356
const STRING = 57357
const BARE_STRING = 57358
const START_TAG = 57359
const CLOSE_TAG = 57360

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"COMMA",
	"DOT",
	"EQ",
	"BRACE_OPEN",
	"BRACE_CLOSE",
	"PAREN_OPEN",
	"PAREN_CLOSE",
	"COMMENT",
	"NUMBER",
	"CRLF",
	"IDENTIFIER",
	"STRING",
	"BARE_STRING",
	"START_TAG",
	"CLOSE_TAG",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:112

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	1, 30,
	6, 21,
	13, 30,
	-2, 17,
}

const yyPrivate = 57344

const yyLast = 56

var yyAct = [...]int{

	15, 3, 19, 38, 54, 19, 36, 35, 51, 17,
	20, 9, 13, 20, 18, 28, 27, 44, 42, 46,
	41, 21, 48, 30, 47, 24, 32, 39, 31, 29,
	25, 52, 49, 40, 16, 26, 12, 11, 10, 8,
	7, 6, 5, 4, 2, 37, 14, 23, 45, 34,
	53, 50, 22, 43, 33, 1,
}
var yyPact = [...]int{

	-2, -1000, 8, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 16, 24, 1, 23, -1000, -1000, -1000,
	-1000, -2, 22, 20, -8, -13, 1, -1000, -1000, 6,
	-1000, 2, 5, 14, 12, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 28, -1000, -1000, -1000, -1000, -1000, -7,
	27, -1000, -11, -1000, -1000,
}
var yyPgo = [...]int{

	0, 55, 54, 53, 52, 51, 50, 49, 48, 47,
	46, 45, 44, 1, 43, 42, 41, 40, 39, 38,
	37, 36, 0, 35, 34, 33,
}
var yyR1 = [...]int{

	0, 1, 12, 12, 13, 13, 13, 13, 13, 13,
	14, 14, 14, 21, 19, 23, 23, 22, 22, 22,
	20, 24, 25, 25, 15, 4, 2, 3, 5, 6,
	16, 17, 9, 7, 8, 18, 10, 11,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 1, 2, 1, 1, 1,
	3, 1, 1, 1, 8, 3, 1, 1, 1, 1,
	1, 4, 3, 1, 1, 3, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -12, -13, -14, -15, -16, -17, -18, 13,
	-19, -20, -21, 14, -10, -22, -24, 11, 16, 4,
	12, 13, -4, -9, 9, 6, -23, -22, 14, 6,
	-13, 6, 6, -2, -7, 15, 14, -11, 16, -22,
	-25, 14, 12, -3, 15, -8, 14, 10, 10, 4,
	-5, 15, 4, -6, 15,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 7, 8, 9,
	10, 11, 12, -2, 0, 0, 0, 13, 36, 18,
	19, 0, 0, 0, 0, 0, 14, 15, 17, 0,
	3, 0, 0, 0, 0, 26, 33, 35, 37, 16,
	20, 22, 23, 0, 27, 31, 34, 25, 32, 0,
	0, 28, 0, 24, 29,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
//line grammar.y:79
		{
			onProject(yyDollar[2].str, yyDollar[4].str, yyDollar[6].str, yyDollar[8].str)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:82
		{
			yyVAL.str = yyDollar[2].str
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:84
		{
			yyVAL.str = yyDollar[1].str
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:86
		{
			yyVAL.str = yyDollar[1].str
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:88
		{
			yyVAL.str = yyDollar[1].str
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:90
		{
			yyVAL.str = yyDollar[1].str
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar.y:95
		{
			onSection(yyDollar[1].str, yyDollar[2].str, yyDollar[4].str)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:98
		{
			yyVAL.str = yyDollar[2].str
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:100
		{
			yyVAL.str = yyDollar[1].str
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:102
		{
			yyVAL.str = yyDollar[1].str
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:105
		{
			onSectionItem(yyDollar[1].str, yyDollar[3].str)
		}
	}
	goto yystack /* stack new state and value */
}
