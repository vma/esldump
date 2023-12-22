# esldump

Connects to freeswitch via ESL and dumps events

## usage

```
Usage: esldump [-h] [-e value] [-H value] [-p value]

 -e, --event=value  events to capture, may be used multiple times (default: CHANNEL_ANSWER
                    CHANNEL_HANGUP)
 -h, --help         print (this) help message
 -H, --host=value   freeswitch ESL host (default: "127.0.0.1")
 -p, --port=value   freeswitch ESL port (default: 8021)
```
