export class MessageLoginRequestV1 {
  email: string;
  password: string;

  constructor() {
    this.email = '';
    this.password = '';
  }

}

export class MessageLoginResponseV1 {
  ok: boolean;
  key: string;

  constructor() {
    this.ok = true;
    this.key = '';
  }

}
