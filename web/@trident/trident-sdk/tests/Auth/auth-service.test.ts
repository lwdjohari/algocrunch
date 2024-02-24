import { Result } from "@trident/trident-core";
import { ClientOptions, ServicePool } from "@trident/trident-core";
import {
    TridentServiceBuilder,
    TridentClientType
} from "../../src/index";

import { RpcError } from '@protobuf-ts/runtime-rpc';

import {
    AuthServiceClientModule as asc,
    AuthServiceModule as aut
} from "../../src/index";

import { AuthService } from "../../src/auth/AuthService";

let authClient: Result<asc.AuthServiceClient, Error>;
let servicePool: ServicePool;
const url = "http://localhost";
const port = 7001;
const wrongUrl = "http://localhost";
const wrongPort = 7002;
let token1: string;
let token2: string;


beforeAll(() => {
    authClient = TridentServiceBuilder
        .getServiceClient<asc.AuthServiceClient>(
            TridentClientType.AUTH_SVC_CLIENT,
            new ClientOptions()
                .setEndpoint(url, port)
                .setFormat('binary'));

    servicePool = new ServicePool("main");
    servicePool.addService<asc.AuthServiceClient>("auth", authClient.unwrap());

});


describe('AuthClient', () => {
    it('should failed the login because wrong credentials', async () => {
        const authClient = servicePool.service<asc.AuthServiceClient>("auth");
        if (authClient.isNone()) {
            throw new Error("AuthClient is not defined");
        }

        const svc = new AuthService(authClient.unwrap());

        const resp = await svc.authenticateUser({
            authType: aut.AuthType.SINGLESTEP,
            password: "",
            persistent: false,
            scope: "",
            username: ""
        });

        if (resp.isOk()) {
            expect(resp.unwrap().status).toBe(401);
        } else {
            // Error happen, usually grpc service cant be found.
            if (resp.error instanceof RpcError) {
                const err = new Error("RpcError: " + resp.error.message + "\nCheck your grpc-server endpoint/addres!");
                err.stack = err.message;
                throw err;
            } else {
                throw resp.error;
            }
        }
    });
});

describe('AuthClient', () => {
    it('should return token if credentials are correct', async () => {
        const authClient = servicePool.service<asc.AuthServiceClient>("auth");
        if (authClient.isNone()) {
            throw new Error("AuthClient is not defined");
        }

        const svc = new AuthService(authClient.unwrap());

        const resp = await svc.authenticateUser({
            authType: aut.AuthType.SINGLESTEP,
            password: "password",
            persistent: false,
            scope: "scope",
            username: "username"
        });

        if (resp.isOk()) {
            expect(resp.unwrap().status).toBe(200);
            expect(resp.unwrap().session!.token).toBeDefined();
        } else {
            // Error happen, usually grpc service cant be found.
            if (resp.error instanceof RpcError) {
                const err = new Error(
                    "RpcError: " +
                    resp.error.message +
                    "\nCheck your grpc-server endpoint/addres!");

                err.stack = err.message;
                throw err;
            } else {
                throw resp.error;
            }
        }
    });
});

describe('AuthClient', () => {
    it('logout should return 404 when no token is found', async () => {
        const authClient = servicePool.service<asc.AuthServiceClient>("auth");
        if (authClient.isNone()) {
            throw new Error("AuthClient is not defined");
        }

        const svc = new AuthService(authClient.unwrap());

        const resp = await svc.logout({ token: "" });

        if (resp.isOk()) {
            expect(resp.unwrap().status).toBe(404);
        } else {
            // Error happen, usually grpc service cant be found.
            if (resp.error instanceof RpcError) {
                const err = new Error("RpcError: " + resp.error.message + "\nCheck your grpc-server endpoint/addres!");
                err.stack = err.message;
                throw err;
            } else {
                throw resp.error;
            }
        }
    });
});

describe('AuthClient', () => {
    it('logout should return 200 when token is exist on server', async () => {
        const authClient = servicePool.service<asc.AuthServiceClient>("auth");
        if (authClient.isNone()) {
            throw new Error("AuthClient is not defined");
        }

        const svc = new AuthService(authClient.unwrap());
        const resp = await svc.logout({ token: "" });

        if (resp.isOk()) {
            expect(resp.unwrap().status).toBe(200);
        } else {
            // Error happen, usually grpc service cant be found.
            if (resp.error instanceof RpcError) {
                const err = new Error("RpcError: " + resp.error.message + "\nCheck your grpc-server endpoint/addres!");
                err.stack = err.message;
                throw err;
            } else {
                throw resp.error;
            }
        }

    });
});


describe('AuthClient', () => {
    it('clearSession should return 404 when is not found', async () => {
        const authClient = servicePool.service<asc.AuthServiceClient>("auth");
        if (authClient.isNone()) {
            throw new Error("AuthClient is not defined");
        }

        const svc = new AuthService(authClient.unwrap());
        const resp = await svc.clearSession({ token: "" });

        if (resp.isOk()) {
            expect(resp.unwrap().status).toBe(404);
        } else {
            // Error happen, usually grpc service cant be found.
            if (resp.error instanceof RpcError) {
                const err = new Error(
                    "RpcError: " +
                    resp.error.message +
                    "\nCheck your grpc-server endpoint/addres!");
                err.stack = err.message;
                throw err;
            } else {
                throw resp.error;
            }
        }
    });
});

describe('AuthClient', () => {
    it('clearSession should return 200 when token is found', async () => {
        const authClient = servicePool.service<asc.AuthServiceClient>("auth");
        if (authClient.isNone()) {
            throw new Error("AuthClient is not defined");
        }

        const svc = new AuthService(authClient.unwrap());
        const resp = await svc.clearSession({ token: "" });

        if (resp.isOk()) {
            expect(resp.unwrap().status).toBe(200);
        } else {
            // Error happen, usually grpc service cant be found.
            if (resp.error instanceof RpcError) {
                const err = new Error(
                    "RpcError: " +
                    resp.error.message +
                    "\nCheck your grpc-server endpoint/addres!");

                err.stack = err.message;
                throw err;
            } else {
                throw resp.error;
            }
        }
    });
});


function expectForErrors(arg: Result<aut.AuthResponse, Error>) {
    expect(arg.isErr()).toBe(true);

    if (arg.isErr()) {
        expect(arg.error).toBeDefined();
    }
};


describe('AuthClient', () => {
    it('handle error when wrong endpoint binding', async () => {
        authClient = TridentServiceBuilder
            .getServiceClient<asc.AuthServiceClient>(
                TridentClientType.AUTH_SVC_CLIENT,
                new ClientOptions()
                    .setEndpoint(wrongUrl, wrongPort)
                    .setFormat('binary'));

        const svc = new AuthService(authClient.unwrap());
        const respCs = await svc.clearSession({ token: "" });
        const respLg = await svc.logout({ token: "" });
        const respAu = await svc.authenticateUser({
            authType: aut.AuthType.SINGLESTEP,
            password: "password",
            persistent: false,
            scope: "scope",
            username: "username"
        });

        expectForErrors(respCs);
        expectForErrors(respLg);
        expectForErrors(respAu);

    });
});

