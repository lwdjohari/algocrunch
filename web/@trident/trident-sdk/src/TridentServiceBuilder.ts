import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { ClientOptions } from "@trident/trident-core";
import {  Err, Ok, Result } from "@trident/trident-core";
import {
  ServiceServiceClientModule,
  UserServiceClientModule,
  AuthServiceClientModule,
  GroupServiceClientModule
} from "../src/index";

export enum TridentClientType {
  UNKNOWN = 0,
  SERVICE_SVC_CLIENT = 1,
  USER_SVC_CLIENT = 2,
  GROUP_SVC_CLIENT = 3,
  AUTH_SVC_CLIENT = 4
}


export class TridentServiceBuilder {

  static getServiceClient<T>(clientType: TridentClientType, arg: ClientOptions): Result<T,Error> {
    let client: T = undefined as T;
    let transport: GrpcWebFetchTransport;
    const is_have_endpoint = arg.endpoint() != null;
    const is_have_credentials = arg.credentials() != null;
    const is_has_additional_headers = arg.headers().size > 0;

    if (!is_have_endpoint) {
      return new Err (new Error('Endpoint is required!'));
    }

    if (is_have_credentials && is_has_additional_headers) {

      transport = new GrpcWebFetchTransport({
        baseUrl: arg.endpoint()!.endpoint(),
        format: arg.format(),
        headers: {
          ...arg.headerAsRecords(),
          [arg.credentials()!.key] :[arg.credentials()!.value]
        }
      });


    } else if (is_have_credentials) {
      transport = new GrpcWebFetchTransport({
        baseUrl: arg.endpoint()!.endpoint(),
        format: arg.format(),
        headers: {
          [arg.credentials()!.key] :[arg.credentials()!.value]
        }
      });


    } else {
      transport = new GrpcWebFetchTransport({
        baseUrl: arg.endpoint()!.endpoint(),
        format: arg.format()
      });

    }

    if (clientType === TridentClientType.SERVICE_SVC_CLIENT) {
      client = new ServiceServiceClientModule.ServicesClient(transport) as T;
    } else if (clientType === TridentClientType.USER_SVC_CLIENT) {
      client = new UserServiceClientModule.UserServiceClient(transport) as T;
    } else if (clientType === TridentClientType.GROUP_SVC_CLIENT) {
      client = new GroupServiceClientModule.GroupServiceClient(transport) as T;
    }else if (clientType === TridentClientType.AUTH_SVC_CLIENT) {
      client = new AuthServiceClientModule.AuthServiceClient(transport) as T;
    }

    return new Ok<T>(client);
  }


}