package com.example.foodapp.ui.auth

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.foodapp.R
import com.example.foodapp.databinding.FragmentLogInBinding
import com.example.foodapp.databinding.FragmentVerificationBinding

class VerificationFragment : Fragment(R.layout.fragment_verification) {
    private  var binding : FragmentVerificationBinding? = null
    private val mBinding get() = binding!!
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentVerificationBinding.inflate(inflater, container, false)
        return mBinding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val nav = findNavController()
        mBinding.imgBackVerification.setOnClickListener {
            nav.popBackStack()
        }
        mBinding.btnNextVerification.setOnClickListener {
            mBinding.progressBar.visibility = View.VISIBLE
        }
    }
}