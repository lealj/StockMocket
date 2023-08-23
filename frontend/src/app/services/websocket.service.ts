import { Injectable } from "@angular/core"
import { Observable, Observer } from 'rxjs';
import { AnonymousSubject, Subject } from "rxjs/internal/Subject";
import { map } from 'rxjs/operators'

const URL = "ws://localhost:8080/ws";

export interface Message {
    source: string;
    content: string;
}

// export class / injectable
@Injectable()
export class WebsocketService {
    // variables
    private subject!: AnonymousSubject<MessageEvent>;
    public messages: Subject<Message>;

    //construct
    constructor() {
        this.messages = <Subject<Message>>this.connect(URL).pipe(
            map(
                (response: MessageEvent): Message => {
                    //console.log('Websocket service msg: ', response.data);
                    let data = JSON.parse(response.data);
                    const message: Message = {
                        source: data.data[0].s,
                        content: data.data[0].p.toString()
                    };
                    console.log(message)
                    return message;
                }
            )
        );
    }

    // functions
    public connect(url: string):AnonymousSubject<MessageEvent> {
        if(!this.subject) {
            this.subject = this.create(url);
            console.log("Connected");
        }
        return this.subject;
    }

    public create(url: string):AnonymousSubject<MessageEvent> {
        let ws = new WebSocket(url);
        let observable = new Observable((obs: Observer<MessageEvent>) => {
            ws.onmessage = obs.next.bind(obs);
            ws.onerror = obs.error.bind(obs);
            ws.onclose = obs.complete.bind(obs);
            return ws.close.bind(obs);
        });
        let observer = {
            error: (err: any)=>{
                console.log('Error in observer: ', err);
            },
            complete: ()=>{},
            next: (data: Object) => {
                console.log('Message sent to websocket: ', data);
                if(ws.readyState === WebSocket.OPEN) {
                    ws.send(JSON.stringify(data));
                }
            }
        };
        return new AnonymousSubject<MessageEvent>(observer, observable);
    }
}
