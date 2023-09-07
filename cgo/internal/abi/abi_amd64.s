#include "go_asm.h"
#include "textflag.h"

#ifdef GOARCH_arm
#define LR R14
#endif

#ifdef GOARCH_amd64
#define	get_tls(r)	MOVQ TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

#ifdef GOARCH_386
#define	get_tls(r)	MOVL TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

/*
    MOVQ	SP, DX
	ANDQ	$~15, SP	// alignment
	MOVQ	DX, 8(SP)
    MOVQ	DI, 48(SP)	// save g
	CALL	R11
	get_tls(CX)
	MOVQ	48(SP), DI
	MOVQ	8(SP), DX
	MOVQ	DX, SP
*/

// expand the stack so that it is C-compatible stack
TEXT ·pushFunc(SB),NOPTR,$1000000-8
    MOVQ fn+0(FP), R11
    RET

// Call C function that returns a float64 with 1 float64 argument 
TEXT ·fcallf(SB),NOSPLIT|NOPTR,$48-16
    MOVQ f1+0(FP), X0

    MOVQ	SP, DX
	ANDQ	$~15, SP	// alignment
	MOVQ	DX, 8(SP)
    MOVQ	DI, 48(SP)	// save g
	CALL	R11
	get_tls(CX)
	MOVQ	48(SP), DI
	MOVQ	8(SP), DX
	MOVQ	DX, SP

    MOVQ X0, ret+8(FP)
    RET
