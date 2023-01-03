<?php

namespace App\Http\Controllers;

use App\Http\Requests\Thematic\StoreThematicRequest;
use App\Http\Resources\ThematicResource;
use App\Models\Thematic;
use Illuminate\Http\Request;

class ThematicController extends Controller
{
    public function index()
    {
        return ThematicResource::collection(Thematic::get());
    }

    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(StoreThematicRequest $request)
    {
        return ThematicResource::make(Thematic::create($request->validated()));
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        return ThematicResource::make(Thematic::where('id', $id)->firstOrFail());
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    public function update(StoreThematicRequest $request, $id)
    {
        return response()->json(Thematic::where('id', $id)->update($request->validated()));
    }

    public function destroy($id)
    {
        return response()->json(Thematic::where('id', $id)->delete());
    }
}
