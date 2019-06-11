export class ItemV1 {
  id: string;
  parent: string;
  title: string;
  body: string;
  createdBy: string;
  waitingFor: string;

  constructor() {
    this.id = '';
    this.parent = '';
    this.title = '';
    this.body = '';
    this.createdBy = '';
    this.waitingFor = '';
  }

}
