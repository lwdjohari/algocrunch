import { Result, Ok, Err,ClientOptions } from "@trident/trident-core";
import { TridentServiceBuilder, TridentClientType } from '../../src/index';

import {
    ServiceServiceClientModule as svc,
    UserServiceClientModule as usc,
    AuthServiceClientModule as asc,
    GroupServiceClientModule as gsc
} from "../../src/index";

describe('Trident AuthClient', () => {
    it('should return error if no endpoint is provided', () => {
        const options = new ClientOptions();
        options.setFormat('binary');

        const client = TridentServiceBuilder
            .getServiceClient<asc.AuthServiceClient>(TridentClientType.AUTH_SVC_CLIENT, new ClientOptions());

        expect(client.isErr()).toBe(true);
    });
});


describe('Trident AuthClient', () => {
    it('should return instance of AuthClient if endpoint is provided', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary');

        const client = TridentServiceBuilder
            .getServiceClient<asc.AuthServiceClient>(TridentClientType.AUTH_SVC_CLIENT, options);

        expect(client.isOk()).toBe(true);
        expect(client.unwrap()).toBeDefined();
    });


});

describe('Trident UserClient', () => {
    it('should return instance of UserClient if endpoint is provided', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary');

        const client = TridentServiceBuilder
            .getServiceClient<usc.UserServiceClient>(TridentClientType.USER_SVC_CLIENT, options);

        expect(client.isOk()).toBe(true);
        expect(client.unwrap()).toBeDefined();
    });
});


describe('Trident ServiceClient', () => {
    it('should return instance of ServicesClient if endpoint is provided', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary');

        const client = TridentServiceBuilder
            .getServiceClient<svc.ServicesClient>(TridentClientType.SERVICE_SVC_CLIENT, options);

        expect(client.isOk()).toBe(true);
        expect(client.unwrap()).toBeDefined();
    });
});

describe('Trident GroupClient', () => {
    it('should return instance of GroupClient if endpoint is provided', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary');

        const client = TridentServiceBuilder
            .getServiceClient<gsc.GroupServiceClient>(TridentClientType.GROUP_SVC_CLIENT, options);

        expect(client.isOk()).toBe(true);
        expect(client.unwrap()).toBeDefined();
    });
});



describe('Trident Get All Client Type', () => {
    it('should return instance of client in TridentClientType', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary');

        const authClient = TridentServiceBuilder
            .getServiceClient(TridentClientType.AUTH_SVC_CLIENT, options);

        const svcClient = TridentServiceBuilder
            .getServiceClient(TridentClientType.SERVICE_SVC_CLIENT, options);

        const userClient = TridentServiceBuilder
            .getServiceClient(TridentClientType.USER_SVC_CLIENT, options);

        const groupClient = TridentServiceBuilder
            .getServiceClient(TridentClientType.GROUP_SVC_CLIENT, options);

        expect(authClient).toBeDefined();
        expect(svcClient).toBeDefined();
        expect(userClient).toBeDefined();
        expect(groupClient).toBeDefined();
    });
});

describe('Trident Client Credential Option Test', () => {
    it('should return instance of client with credentials', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary')
            .setCredentials('token');

        const authClient = TridentServiceBuilder
            .getServiceClient<asc.AuthServiceClient>(TridentClientType.AUTH_SVC_CLIENT, options);

        expect(authClient.isOk()).toBe(true);
        expect(authClient.unwrap()).toBeDefined();
    });
});


describe('Trident Client All Options Test', () => {
    it('should return instance of client with all options set', () => {
        const options = new ClientOptions()
            .setEndpoint("http:localhost")
            .setFormat('binary')
            .setCredentials('token')
            .addHeaders('custom-header-1', 'val')
            .addHeaders('custom-header-2', 'val2');

        const authClient = TridentServiceBuilder
            .getServiceClient<asc.AuthServiceClient>(TridentClientType.AUTH_SVC_CLIENT, options);

        expect(authClient.isOk()).toBe(true);
        expect(authClient.unwrap()).toBeDefined();
    });
});


