import { Option, Some, None, Result, Ok, Err } from '@trident/trident-core';
import { ITridentService } from '../ITridentService';
import { TridentServiceBuilder, TridentClientType } from '../TridentServiceBuilder';
import { ClientOptions } from '@trident/trident-core';

import {
    UserServiceClientModule as usc,
    UserServiceModule as usm,
    CommonModule as cmn
} from "../trident";
import { Timestamp } from '../trident/google/protobuf/timestamp';

export class UserCredentialService implements ITridentService{
    private client_: Option<usc.UserServiceClient> = new None();

    constructor(args: ClientOptions | usc.UserServiceClient) {
        if (args instanceof usc.UserServiceClient) {
            this.client_ = new Some(args);
        } else {
            const client = TridentServiceBuilder
                .getServiceClient<usc.UserServiceClient>(
                    TridentClientType.USER_SVC_CLIENT,
                    args);

            this.client_ = new Some(client.unwrap());
        }
    }

    public getClient(): Option<usc.UserServiceClient> {
        return this.client_;
    }
}