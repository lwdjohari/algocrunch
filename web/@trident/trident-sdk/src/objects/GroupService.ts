import { Option, Some, None, Result, Ok, Err } from '@trident/trident-core';
import { ITridentService } from '../ITridentService';
import { TridentServiceBuilder, TridentClientType } from '../TridentServiceBuilder';
import { ClientOptions } from '@trident/trident-core';

import {
    GroupServiceClientModule as gsc,
    GroupServiceModule as gsm,
    CommonModule as cmn
} from "../trident";
import { Timestamp } from '../trident/google/protobuf/timestamp';

export class GroupService implements ITridentService{

    private client_: Option<gsc.GroupServiceClient> = new None();

    public constructor(args: ClientOptions | gsc.GroupServiceClient) {
        if (args instanceof gsc.GroupServiceClient) {
            this.client_ = new Some(args);
        } else {
            const client = TridentServiceBuilder
                .getServiceClient<gsc.GroupServiceClient>(
                    TridentClientType.GROUP_SVC_CLIENT,
                    args);

            this.client_ = new Some(client.unwrap());
        }
    }

    public getClient(): Option<gsc.GroupServiceClient> {
        return this.client_;
    }

    
    
}