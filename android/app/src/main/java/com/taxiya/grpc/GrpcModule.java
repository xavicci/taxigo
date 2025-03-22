package com.taxiya.grpc;

import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.ReadableMap;
import com.facebook.react.bridge.WritableMap;
import com.facebook.react.bridge.Arguments;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import taxiya.proto.auth.AuthServiceGrpc;
import taxiya.proto.auth.LoginRequest;
import taxiya.proto.auth.RegisterRequest;
import taxiya.proto.auth.LoginResponse;
import taxiya.proto.auth.RegisterResponse;

public class GrpcModule extends ReactContextBaseJavaModule {
    private final ReactApplicationContext reactContext;
    private ManagedChannel channel;
    private AuthServiceGrpc.AuthServiceStub authStub;

    public GrpcModule(ReactApplicationContext reactContext) {
        super(reactContext);
        this.reactContext = reactContext;
        setupGrpc();
    }

    private void setupGrpc() {
        channel = ManagedChannelBuilder.forAddress("10.0.2.2", 50051)
                .usePlaintext()
                .build();
        authStub = AuthServiceGrpc.newStub(channel);
    }

    @Override
    public String getName() {
        return "GrpcModule";
    }

    @ReactMethod
    public void login(ReadableMap request, Promise promise) {
        LoginRequest loginRequest = LoginRequest.newBuilder()
                .setEmail(request.getString("email"))
                .setPassword(request.getString("password"))
                .build();

        authStub.login(loginRequest, new StreamObserver<LoginResponse>() {
            @Override
            public void onNext(LoginResponse response) {
                WritableMap result = Arguments.createMap();
                result.putString("token", response.getToken());
                result.putMap("user", Arguments.createMap());
                promise.resolve(result);
            }

            @Override
            public void onError(Throwable t) {
                promise.reject("GRPC_ERROR", t.getMessage());
            }

            @Override
            public void onCompleted() {
                // No-op
            }
        });
    }

    @ReactMethod
    public void register(ReadableMap request, Promise promise) {
        RegisterRequest registerRequest = RegisterRequest.newBuilder()
                .setEmail(request.getString("email"))
                .setPassword(request.getString("password"))
                .setName(request.getString("name"))
                .setPhone(request.getString("phone"))
                .build();

        authStub.register(registerRequest, new StreamObserver<RegisterResponse>() {
            @Override
            public void onNext(RegisterResponse response) {
                WritableMap result = Arguments.createMap();
                result.putString("token", response.getToken());
                result.putMap("user", Arguments.createMap());
                promise.resolve(result);
            }

            @Override
            public void onError(Throwable t) {
                promise.reject("GRPC_ERROR", t.getMessage());
            }

            @Override
            public void onCompleted() {
                // No-op
            }
        });
    }
} 