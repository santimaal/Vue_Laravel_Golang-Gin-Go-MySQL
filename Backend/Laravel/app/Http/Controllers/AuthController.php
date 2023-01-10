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

        if ($user) {
            return response()->json('User created successfully', 200);
        }else {
            return response()->json('Error Register', 401);
        }
    }


    //Login User
    public function login(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'email' => 'required|email',
            'password' => 'required|string|min:2',
        ]);

        if ($validator->fails()) {
            return response()->json($validator->errors(), 422);
        } else {
            $user_exist = User::select()->where('email', $request->email)->get()->count();
            if ($user_exist) {
                $credentials = $request->only('email', 'password');
                $token = Auth::attempt($credentials);
                if (!$token = auth()->attempt($validator->validated())) {
                    return response()->json(['error' => 'Unauthorized'], 401);
                } else {
                    if (auth()->user()->type == "admin") {
                        $user_serialize = Auth::user();
                        unset($user_serialize->password);
                        return response()->json([
                            'status' => 'success',
                            'user' => $user_serialize,
                            'authorisation' => [
                                'token' => $token,
                                'type' => $user_serialize->type,
                            ]
                        ]);
                    } else {
                        return response()->json(['error' => 'Unauthorized'], 401);
                    }
                }
            } else {
                return response()->json(['error' => 'User not exist'], 401);
            }
        }
    }

    //PROFIILE USER
    public function GetProfile(){
        $user_serialize = Auth::user();
        unset($user_serialize->password);
        return response()->json($user_serialize);
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
