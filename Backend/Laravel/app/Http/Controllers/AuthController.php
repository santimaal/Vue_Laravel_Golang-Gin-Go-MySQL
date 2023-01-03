<?php

namespace App\Http\Controllers;

use App\Http\Requests\User\StoreUserRequest;
use App\Http\Requests\User\RegisterUserRequest;
use App\Http\Resources\UserResource;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use App\Models\User;


class AuthController extends Controller
{
    protected User $user;

    public function __construct(User $user)
    {
        $this->user = $user;
    }

    //Register USER
    public function register(RegisterUserRequest $request)
    {
        $user = User::create([
            'name' => $request->name,
            'email' => $request->email,
            'password' => Hash::make($request->password),
        ]);

        $token = Auth::login($user);
        $type= User::select('type')->where('id', $user['id'] )->get()[0]['type'];
        return response()->json([
            'status' => 'success',
            'message' => 'User created successfully',
            'user' => $user,
            'authorisation' => [
                'token' => $token,
                'type' => $type,
            ]
        ]);
    }


    //Login User
    public function login(Request $request)
    {

        $validator = Validator::make($request->all(), [
            'email' => 'required|email',
            'password' => 'required|string|min:6',
        ]);

        if ($validator->fails()) {
            echo ($validator->errors());
            return response()->json($validator->errors(), 422);
        } else {
            echo "Soy el ELSE";
        }



        // if (! $token = auth()->attempt($validator->validated())) {
        //     return response()->json(['error' => 'Unauthorized'], 401);
        // }

        // return $this->createNewToken($token);



        // $request->validate([
        //     'email' => 'required|string|email',
        //     'password' => 'required|string|min:6',
        // ]);

        // $credentials = $request->only('email', 'password');
        // $token = Auth::attempt($credentials);


        // if (!$token) {
        //     return response()->json([
        //         'status' => 'error',
        //         'message' => 'Unauthorized',
        //     ], 401);
        // }else{
        //     echo "EXISTE TOKEN";
        // }

        // $user = Auth::user();
        // return response()->json([
        //     'status' => 'success',
        //     'user' => $user,
        //     'authorisation' => [
        //         'token' => $token,
        //         'type' => 'bearer',
        //     ]
        // ]);
    }

    // GET ALL USER
    public function index()
    {
        return UserResource::collection(User::get());
    }

    //GET ONE USER
    public function show($id)
    {
        return UserResource::make(User::where('id', $id)->firstOrFail());
    }

    // CREATE A USER
    public function store(StoreUserRequest $request)
    {
        return UserResource::make(User::create($request->validated()));
    }

    //UPDATED A USER
    public function update(Request $request, $id)
    {
        return response()->json(User::where('id', $id)->update($request->validated()));
    }

    // DELETE USER 
    public function destroy($id)
    {
        return response()->json(User::where('id', $id)->delete());
    }
}
