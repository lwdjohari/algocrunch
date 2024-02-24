import { Option, Some, None, Result, Ok, Err } from '@trident/trident-core';
import { ITridentService } from '../ITridentService';
import { TridentServiceBuilder, TridentClientType } from '../TridentServiceBuilder';
import { ClientOptions } from '@trident/trident-core';
import {
    UserServiceClientModule as usc,
    UserServiceModule as usr,
    CommonModule as cmn
} from "../trident";
import { Timestamp } from '../trident/google/protobuf/timestamp';


export class UserService implements ITridentService {

    private client_: Option<usc.UserServiceClient> = new None();

    public constructor(args: ClientOptions | usc.UserServiceClient) {
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

    public createUserObject(args: {
        userId: Option<string>,
        username: string,
        password: string,
        email: string,
        identityProvider: usr.IdentityProvider,
        identityType: usr.IdentityType,
        status: usr.UserStatus
    }): Result<usr.User, Error> {

        const now = new Date();
        const object = usr.User.create({
            userId: args.userId.isSome() ? args.userId.unwrap() : undefined,
            username: args.username,
            password: args.password,
            email: args.email,
            identityProvider: args.identityProvider,
            identityType: args.identityType,
            status: args.status,
            createdOn: Timestamp.fromDate(now),
            updatedOn: Timestamp.fromDate(now)
        });

        return new Ok(object);
    }

    public createUserExtInfoObject(args: {
        userId: Option<string>,
        userExtendedInfoId: Option<string>,
        type: number,
        key: string,
        caption: string,
        value: string
    }): Result<usr.UserExtendedInfo, Error> {
        const now = new Date();
        const object = usr.UserExtendedInfo.create(
            {
                userId: args.userId.isSome() ? args.userId.unwrap() : undefined,
                userExtendedInfoId: args.userId.isSome() ? args.userExtendedInfoId.unwrap() : undefined,
                type: args.type,
                key: args.key,
                caption: args.caption,
                value: args.value,
                createdOn: Timestamp.fromDate(now),
                updatedOn: Timestamp.fromDate(now)
            }
        );


        return new Ok(object);
    }

    public createUserProfileObject(args: {

    }): Result<usr.UserProfile, Error> {

        const now = new Date();
        const object = usr.UserProfile.create({
            
        });

        return new Ok(object)
    }

    public createAvatarObject(args: {

    }): Result<usr.Avatar, Error> {

        const now = new Date();
        const object = usr.Avatar.create({

        });

        return new Ok(object);
    }


    public async addNewUser(
        scope: string,
        user: usr.User,
        profile: Option<usr.UserProfile>,
        extInfo: Option<usr.UserExtendedInfo>,
        avatar: Option<usr.Avatar>,
        actions: usr.UserServiceWriteAction
    ): Promise<Result<usr.UserServiceWriteResponse, Error>> {

        let modules = usr.UserServiceModule.USER;
        modules = profile.isSome() ? modules | usr.UserServiceModule.USER_PROFILE : modules;
        modules = avatar.isSome() ? modules | usr.UserServiceModule.USER_AVATAR : modules;
        modules = extInfo.isSome() ? modules | usr.UserServiceModule.USER_EXTENDED_INFO : modules;

        const request = usr.UserServiceWriteRequest.create({
            action: usr.UserServiceWriteAction.USER_SERVICE_ACTION_CREATE,
            module: modules,
            user: user,
            userAvatar: avatar.isSome() ? avatar.unwrap() : undefined,
            userExtendedInfo: extInfo.isSome() ? extInfo.unwrap() : undefined,
            userProfile: profile.isSome() ? profile.unwrap() : undefined
        });

        try {
            const response = await this.client_.unwrap().createUser(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    }

    public async updateUser(
        scope: string,
        user: usr.User,
        profile: Option<usr.UserProfile>,
        extInfo: Option<usr.UserExtendedInfo>,
        avatar: Option<usr.Avatar>
    ): Promise<Result<usr.UserServiceWriteResponse, Error>> {

        let modules = usr.UserServiceModule.USER;
        modules = profile.isSome() ? modules | usr.UserServiceModule.USER_PROFILE : modules;
        modules = avatar.isSome() ? modules | usr.UserServiceModule.USER_AVATAR : modules;
        modules = extInfo.isSome() ? modules | usr.UserServiceModule.USER_EXTENDED_INFO : modules;

        const request = usr.UserServiceWriteRequest.create({
            action: usr.UserServiceWriteAction.USER_SERVICE_ACTION_UPDATE,
            module: modules,
            user: user,
            userAvatar: avatar.isSome() ? avatar.unwrap() : undefined,
            userExtendedInfo: extInfo.isSome() ? extInfo.unwrap() : undefined,
            userProfile: profile.isSome() ? profile.unwrap() : undefined
        });

        try {
            const response = await this.client_.unwrap().updateUser(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    }

    public async requestForUserDeletion(
        scope: string,
        userId: string
    ): Promise<Result<usr.UserServiceWriteResponse, Error>> {

        const request = usr.UserServiceWriteRequest.create({
            action: usr.UserServiceWriteAction.USER_SERVICE_ACTION_DELETE,
            module: usr.UserServiceModule.USER,
            user: usr.User.create({ userId: userId })
        });

        try {
            const response = await this.client_.unwrap().deleteUser(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    }

    public async deleteUser(
        scope: string,
        userId: string
    ): Promise<Result<usr.UserServiceWriteResponse, Error>> {

        const request = usr.UserServiceWriteRequest.create({
            action: usr.UserServiceWriteAction.USER_SERVICE_ACTION_DELETE,
            module: usr.UserServiceModule.USER,
            user: usr.User.create({ userId: userId })
        });

        try {
            const response = await this.client_.unwrap().deleteUser(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    };

    public async getUser(
        scope: string,
        userId: string,
        modules: usr.UserServiceModule
    ): Promise<Result<usr.UserServiceListResponse, Error>> {

        modules = modules | usr.UserServiceModule.USER;

        let query = usr.UserServiceQuery.create({
            userIds: [userId]
        });

        const request = usr.UserServiceListRequest.create({
            module: modules,
            scope: scope,
            filters: query
        });

        try {
            const response = await this.client_.unwrap().getUser(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    }

    public async getUsers(
        scope: string,
        query: Option<string>,
        filters: Option<usr.UserServiceQuery>,
        modules: usr.UserServiceModule,
        page: Option<cmn.PageRequest>
    ): Promise<Result<usr.UserServiceListResponse, Error>> {

        modules = modules | usr.UserServiceModule.USER;

        const request = usr.UserServiceListRequest.create({
            module: modules,
            scope: scope,
            filters: filters.isSome() ? filters.unwrap() : undefined,
            query: query.isSome() ? query.unwrap() : undefined,
            page: page.isSome() ? page.unwrap() : undefined
        });

        try {
            const response = await this.client_.unwrap().listUsers(request);
            return new Ok(response.response);
        } catch (error) {
            return new Err(error as Error);
        }
    }

    public getClient(): Option<usc.UserServiceClient> {
        return this.client_;
    }
}

