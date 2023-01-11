<?php

namespace App\Http\Controllers;

use App\Http\Requests\Reserve\StoreReserveRequest;
use App\Http\Resources\ReserveResource;
use App\Models\Reserve;
use Illuminate\Http\Request;

class ReserveController extends Controller
{
    public function index()
    {
        return ReserveResource::collection(Reserve::where('is_confirmed', "pending")->get());
    }

    public function create()
    {
        //
    }

   
    public function store(StoreReserveRequest $request)
    {
        return ReserveResource::make(Reserve::create($request->validated()));
    }

 
    public function show($id)
    {
        return ReserveResource::make(Reserve::where('id', $id)->firstOrFail());
    }

   
    public function edit($id)
    {
        //
    }

    public function update(Request $request, $id)
    {
        return response()->json(Reserve::where('id', $id)->update($request->validated()));
    }

    public function destroy($id)
    {
        return response()->json(Reserve::where('id', $id)->delete());
    }
}