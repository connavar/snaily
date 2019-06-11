import {ItemV1} from './item-v1';
import {UserV1} from './user-v1';

export class MyItemsRequestV1 {
  createdByMe: ItemV1[];
  waitingForMe: ItemV1[];
  users: UserV1[];

  constructor() {
    this.createdByMe = [];
    this.waitingForMe = [];
    this.users = [];
  }

}


