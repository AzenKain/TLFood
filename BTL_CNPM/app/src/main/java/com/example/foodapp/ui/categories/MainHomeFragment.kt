package com.example.foodapp.ui.categories

import android.content.Intent
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import androidx.recyclerview.widget.GridLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.foodapp.ProductDetailActivity
import com.example.foodapp.R
import com.example.foodapp.adapter.ProductAdapter
import com.example.foodapp.data.Product
import com.example.foodapp.databinding.FragmentMainHomeBinding

class MainHomeFragment : Fragment() {
    private lateinit var binding:FragmentMainHomeBinding
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentMainHomeBinding.inflate(layoutInflater)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val ds = mutableListOf<Product>()
        ds.add(Product(R.drawable.pizza1, "Monkey.D Luffy", 21211111, 20000))
        ds.add(Product(R.drawable.s4, "One Piece", 2000000, 254))
        ds.add(Product(R.drawable.fries3, "Zoro", 2000000, 1240))
        ds.add(Product(R.drawable.pizza1, "Monkey.D Luffy", 21211111, 1))
        ds.add(Product(R.drawable.s4, "One Piece", 2000000, 0))
        ds.add(Product(R.drawable.fries3, "Zoro", 2000000, 5))
        val adapter = ProductAdapter(ds){ product ->
            val intent = Intent(requireContext(), ProductDetailActivity::class.java)
            intent.putExtra("product", product)
            startActivity(intent)

        }
        val rcv = view.findViewById<RecyclerView>(R.id.recyclerViewProducts1)
        rcv.adapter = adapter
        rcv.layoutManager = GridLayoutManager(context, 1)
        binding.progressBar.visibility = View.GONE

    }
}