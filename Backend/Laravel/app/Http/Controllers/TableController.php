<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreTableRequest;
use App\Http\Resources\TableResource;
use App\Models\Table;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class TableController extends Controller
{

    // GETALL
    public function index()
    {
        return TableResource::collection(Table::get());
    }

    public function create()
    {
        //
    }

    // CREATE
    public function store(StoreTableRequest $request)
    {
        return TableResource::make(Table::create($request->validated()));
    }

    // GETONE
    public function show($id)
    {
        return TableResource::make(Table::where('id', $id)->firstOrFail());
    }

    public function edit($id)
    {
        //
    }

    // UPDATE
    public function update(StoreTableRequest $request, $id)
    {
        return response()->json(Table::where('id', $id)->update($request->validated()));
    }

    // DELETE
    public function destroy($id)
    {
        return response()->json(Table::where('id', $id)->delete());
    }
}
