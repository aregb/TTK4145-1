# TTK4145

Sanntidsprogrammering NTNU -> aka verdens beste lab <3

Problemer/issues:
* noen ganger når en heis er på så tar den CAB order to ganger -> altså den kommer til etg. stopper, lukker opp dører starter timer -> SÅ lukker den dører-> åpner dører igjen og først DA skrur den av lyset i CAB button..
* alle heiser på nettverket tar alle ordre(hvertfall CAB orders) og de andre fordeles trolig feil -> alt dette kommer fra feil FØR det sendes inn til .D filen. Altså har alle tilordnet seg selv CAB ordren før den filen kalles
* lagt opp printscreen som viser en issue, med at begge tar samme "cab" order. Dette er printet i Consensus cab og det er at begge tror de skal ha den cab orderen og at den andre ikke har noen cab order


Usefull CMDs:
* KJØR: "go build main.go && ./main --id=123
* "who" -> finner ut om noen er SSHet seg på, set etter "ipen" på slutten
* "ps -aux|grep < id fra who >" e.g. pts/0
* "kill -9 <prosess id fra sshd> (ene linje fra over)" e.g. 7427

