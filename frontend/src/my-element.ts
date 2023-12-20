import { css, html, LitElement } from "lit";
import { customElement, property } from "lit/decorators.js";
import { Connect } from "../wailsjs/go/main/App";
import * as Runtime from "../wailsjs/runtime/runtime.js";
import "./style.css";

/**
 * An example element.
 *
 * @slot - This element has a slot
 * @csspart button - The button
 */
@customElement("my-element")
export class MyElement extends LitElement {
  static styles = css`
    #logo {
      display: block;
      width: 50%;
      height: 50%;
      margin: auto;
      padding: 10% 0 0;
      background-position: center;
      background-repeat: no-repeat;
      background-size: 100% 100%;
      background-origin: content-box;
    }

    .result {
      height: 20px;
      line-height: 20px;
      margin: 1.5rem auto;
    }

    .input-box .btn {
      width: 60px;
      height: 30px;
      line-height: 30px;
      border-radius: 3px;
      border: none;
      margin: 0 0 0 20px;
      padding: 0 8px;
      cursor: pointer;
    }

    .input-box .btn:hover {
      background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
      color: #333333;
    }

    .input-box .input {
      border: none;
      border-radius: 3px;
      outline: none;
      height: 30px;
      line-height: 30px;
      padding: 0 10px;
      background-color: rgba(240, 240, 240, 1);
      -webkit-font-smoothing: antialiased;
    }

    .input-box .input:hover {
      border: none;
      background-color: rgba(255, 255, 255, 1);
    }

    .input-box .input:focus {
      border: none;
      background-color: rgba(255, 255, 255, 1);
    }
  `;

  @property()
  speed = 0;

  connect() {
    Connect().then();
  }

  render() {
    return html`
      <main>
        <div class="result" id="speed">${this.speed}</div>
        <div class="input-box" id="input">
          <button @click=${this.connect} class="btn">Connect</button>
        </div>
      </main>
    `;
  }

  connectedCallback(): void {
    super.connectedCallback();
    Runtime.EventsOn("speed", (speed) => {
      this.speed = speed;
    });
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "my-element": MyElement;
  }
  interface Window {
    runtime: {
      EventsOn: (
        eventName: string,
        callback: (optionalData?: any) => void
      ) => () => void;
      EventsOff: (eventName: string, ...additionalEventNames: string[]) => void;
      EventsOnce: (
        eventName: string,
        callback: (optionalData?: any) => void
      ) => () => void;
      EventsOnMultiple: (
        eventName: string,
        callback: (optionalData?: any) => void,
        counter: number
      ) => () => void;
      EventsEmit: (eventName: string, ...optionalData: any) => () => void;
    };
  }
}
