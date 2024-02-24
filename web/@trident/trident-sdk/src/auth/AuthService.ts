import { ClientOptions } from "@trident/trident-core";
import { AuthServiceClientModule as asc, AuthServiceModule as aut } from "../trident";
import { TridentServiceBuilder, TridentClientType } from "../TridentServiceBuilder";
import { Err, None, Ok, Option, Result, Some } from "@trident/trident-core";
import { ITridentService } from "../ITridentService";


export class AuthService implements ITridentService{

    private client_: Option<asc.AuthServiceClient> = new None();

    public constructor(args: ClientOptions | asc.AuthServiceClient) {
        if (args instanceof asc.AuthServiceClient) {
            this.client_ = new Some(args);
        } else {
            const rest = TridentServiceBuilder
                .getServiceClient<asc.AuthServiceClient>(
                    TridentClientType.AUTH_SVC_CLIENT,
                    args);

            this.client_ = new Some(rest.unwrap());
        }
    }

    public async authenticateUser(args: {
        username: string, password: string,
        scope: string, persistent: boolean,
        authType: aut.AuthType
    }): Promise<Result<aut.AuthResponse,Error>> {
        const request = aut.Auth.create({
            flags: aut.AuthFlags.AUTH_FLAG_USERNAME |
                aut.AuthFlags.AUTH_FLAG_PASS |
                aut.AuthFlags.AUTH_FLAG_SCOPE |
                aut.AuthFlags.AUTH_FLAG_PERSISTENT,
            password: args.password,
            scope: args.scope,
            username: args.username,
            persistent: args.persistent,
            type: args.authType
        });

        try {
            const response = await this.client_.unwrap().authenticate(request);
            return new Ok(response.response);
        }catch(error) {
            return new Err( error as Error);
        }
        
    }

    public getClient(): Option<asc.AuthServiceClient> {
        return this.client_!;
    }

    public async logout(args: { token: string }): Promise<Result<aut.AuthResponse,Error>> {
        const request = aut.AuthSession.create({
            token: args.token
        });

        try {
            const response = await this.client_.unwrap().logout(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }

    }

    public async clearSession(args: { token: string }): Promise<Result<aut.AuthResponse,Error>> {
        const request = aut.AuthSession.create({
            token: args.token
        });

        try {
            const response = await this.client_.unwrap().clearSession(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    }

    public dispose() {
        if (this.client_.isSome()) {
            this.client_ = new None();
        }
    }
}
