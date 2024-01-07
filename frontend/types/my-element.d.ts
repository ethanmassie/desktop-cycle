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
}
