# BRIQ-CLI

### TO START
 ```
 > git clone https://github.com/nlm/briq-cli.git
 > cd briq-cli
 > make
 > ./briq-cli -h
 ```

### CONFIG

- copy the `config.example.yaml` and name it `config.yaml`
- find your token by inspecting request in briq website (in the `Authorization` header)
- copy your token and paste it in the `config.yaml`
- create as much **groups** as you want by adding the usernames of the people you want to send briq to

### COMMANDS

- `give`, `g`

    Give one or more briqs to someone
    ```
    > briq-cli give [flags]
    ```
    | Flags |  Type | Default value| Explanation | 
    | ------ | :--: | :--: | ------ |
    | --amount | uint | 1 | How Many briqs to give to the user |
    | --message | string | "Have a Briq! #Rock-solid"| Message to send, must include one #Value |
    | --to | string | | Username to give to |

- `group-give`, `gg`

    Give one or more briqs to a **group**
    ```
    > briq-cli group-give [flags]
    ```

    | Flags |  Type | Default value| Explanation | 
    | ------ | :--: | :--: | ------ |
    | --amount | uint | 1 | How Many briqs to give to each user |
    | --message | string | "🎁 #Community"| Message to send, must include one #Value |
    | --to-group | string | | Limit to a specific group of users |

- `help`

    Help provides help for any command in the application.
    ```
    > briq-cli help [path to command]
    ```

- `list-users`, `lu`

    List info about existing users on briq 
    ```
    > ./briq-cli lu
    +---------------------------+--------------------------+------+---------+---------+
    | USERNAME                  | DISPLAY NAME             |   XP | BALANCE | TO GIVE |
    +---------------------------+--------------------------+------+---------+---------+
    | jdupont                   | Jean Dupont              |  42  |     123 |      12 |
    ...
    ```
- `me`, `m`

    Get information about me
    ```
    > ./briq-cli me
    +-----------------+--------------------------------------+
    | KEY             | VALUE                                |
    +-----------------+--------------------------------------+
    | Id              | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx |
    | Username        | jdupont                              |
    | Full Name       | Jean Dupont                          |
    | Display Name    | Jean                                 |
    | Briqs to give   | 12                                   |
    | Collected Briqs | 123                                  |
    +-----------------+--------------------------------------+
    ```

- `random-love`, `rl`

    Send a briq to a number of random users
    ```
    > briq-cli random-love [flags]
    ```
    | Flags |  Type | Default value| Explanation | 
    | ------ | :--: | :--: | ------ |
    | --amount | uint | 1 | How Many briqs to give to each user |
    | --message | string | "Everybody needs some Briqs #Community"| Message to send, must include one #Value |
    | --to-group | string | | Limit to a specific group of users |
    | --user-count | uint | 3 | Number of users to send briqs to |



### NOTES:
- You can only send a maximum of 4 briq to the same user per day
- You can use `-h`, `--help` for each command to get help

