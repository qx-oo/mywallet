# mywallet

This is command line eth wallet

### Command

Create new wallet:

    create -pass PASSWORD -name NAME

Show all wallet:

    show

Delete wallet:

    delete -pass PASSWROD -name NAME

Import wallet by mnemonic:

    import -words "xx xx xx ... "

Query account balance:

    balance -pass PASSWROD -name NAME

Transfer:

    transfer -pass PASSWORD -name NAME -to TOADDR -value VALUE

### Run

    mkdir datadir
    go run main.go <cmd> <args>