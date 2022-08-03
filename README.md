# mywallet

This is command line eth wallet

### Init

ABI File:

    solc --abi sol/Mytoken.sol -o sol

Sol Bin:

    solc --bin sol/Mytoken.sol -o sol

Go ABI:

    abigen -abi sol/Mytoken.abi -pkg sol -out sol/mytoken.go

Go ABI(Deploy):

    abigen --bin=sol/Mytoken.bin --abi=sol/Mytoken.abi --pkg=sol --out=sol/mytoken.go

### Command

Command

	create -pass PASSWORD -name NAME --for create new wallet
	show --for show all wallet
	delete -pass PASSWROD -name NAME --for delete wallet
	import -words "xx xx xx ... " --for import wallet by mnemonic
	balance -pass PASSWROD -name NAME --for query account balance
	transfer -pass PASSWORD -name NAME -to TOADDR -value VALUE --for transfer from acct to toaddr

Token Command

	deploytoken -pass PASSWORD -name NAME --for deploy token
	minttoken -pass PASSWORD -name NAME -to TOADDR -value VALUE --for mint token to toaddr
	sendtoken -pass PASSWORD -name NAME -to TOADDR -value VALUE --for send mytoken
	balancetoken -name NAME --for query account balance

### Run

    mkdir datadir
    go run main.go <cmd> <args>