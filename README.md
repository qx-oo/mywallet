# mywallet

### Command

    create -pass PASSWORD -name NAME --for create new wallet
    show --for show all wallet
    delete -pass PASSWROD -name NAME --for delete wallet
    import -words "xx xx xx ... " --for import wallet by mnemonic
    balance -pass PASSWROD -name NAME --for query account balance
    transfer -pass PASSWORD -name NAME -to TOADDR -value VALUE --for transfer from acct to toaddr

### Run

    mkdir datadir
    go run main.go <cmd> <args>