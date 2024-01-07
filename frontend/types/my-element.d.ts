import { LitElement } from "lit";
import "./style.css";
/**
 * An example element.
 *
 * @slot - This element has a slot
 * @csspart button - The button
 */
export declare class MyElement extends LitElement {
    static styles: import("lit").CSSResult;
    speed: number;
    connect(): void;
    render(): import("lit-html").TemplateResult<1>;
    connectedCallback(): void;
}
declare global {
    interface HTMLElementTagNameMap {
        "my-element": MyElement;
    }
    interface Window {
        runtime: {
            EventsOn: (eventName: string, callback: (optionalData?: any) => void) => () => void;
            EventsOff: (eventName: string, ...additionalEventNames: string[]) => void;
            EventsOnce: (eventName: string, callback: (optionalData?: any) => void) => () => void;
            EventsOnMultiple: (eventName: string, callback: (optionalData?: any) => void, counter: number) => () => void;
            EventsEmit: (eventName: string, ...optionalData: any) => () => void;
        };
    }
}
