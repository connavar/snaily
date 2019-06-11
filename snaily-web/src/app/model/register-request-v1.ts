export class MessageRegisterRequestV1 {
  email: string;
  password: string;
  firstname: string;
  lastname: string;

  constructor() {
    this.email = '';
    this.password = '';
    this.firstname = '';
    this.lastname = '';
  }

}

export class MessageRegisterResponseV1 {
  ok: boolean;
}
