<?php

namespace App\Http\Controllers;

use App\Http\Requests\Reserve\StoreReserveRequest;
use App\Http\Resources\ReserveResource;
use App\Models\Reserve;
use Illuminate\Support\Facades\DB;

class ReserveController extends Controller
{
    public function index()
    {
        // return response()->json(Reserve::join('users', 'reserves.id_user', '=', 'users.id')->where('reserves.is_confirmed',"pending")->select('reserves.*', 'users.name')->get());
        return response()->json(DB::select("CALL getReservesAdmin()"));
    }

    public function store(StoreReserveRequest $request)
    {
        return ReserveResource::make(Reserve::create($request->validated()));
    }


    public function show($id)
    {
        return ReserveResource::make(Reserve::where('id', $id)->firstOrFail());
    }

    public function update($id, StoreReserveRequest $request)
    {
        return response()->json(Reserve::where('id', $id)->update($request->validated()));
    }

    public function destroy($id)
    {
        return response()->json(Reserve::where('id', $id)->delete());
    }
}
